package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/johannes94/blog/posts/crud-rest-api-with-gorilla-mux/middleware/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        string `json:"id"`
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
}

func dropTable(db *gorm.DB) error {
	return db.Exec("DROP TABLE users;").Error
}

func initUsers(db *gorm.DB) error {
	data, err := ioutil.ReadFile("userdata.json")
	if err != nil {
		return err
	}

	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		return err
	}

	return db.Create(&users).Error
}

func handleFlags(db *gorm.DB) {
	initData := flag.Bool("init-data", false, "Set this flag if DB should be initialized with dummy data")
	drop := flag.Bool("drop-table", false, "Set this flag if you wan't to drop all user data in your DB")
	flag.Parse()

	if *drop {
		msg := ""
		if err := dropTable(db); err != nil {
			msg = fmt.Sprintf("Error dropping table: %v", err)
		} else {
			msg = "Dropped users table in DB"
		}

		fmt.Println(msg)
		os.Exit(0)
	}

	if *initData {
		msg := ""
		if err := initUsers(db); err != nil {
			msg = fmt.Sprintf("Error initializing data in DB: %v", err)
		} else {
			msg = "Initialized data in DB."
		}

		fmt.Println(msg)
		os.Exit(0)
	}
}

func main() {

	db, err := initDB()
	if err != nil {
		fmt.Printf("Error initializing DB: %v\n", err)
	}

	handleFlags(db)

	uh := userHandler{db: db}

	r := mux.NewRouter()
	usersR := r.PathPrefix("/users").Subrouter()
	usersR.Path("").Methods(http.MethodGet).HandlerFunc(uh.getAllUsers)
	usersR.Path("").Methods(http.MethodPost).HandlerFunc(uh.createUser)
	usersR.Path("/{id}").Methods(http.MethodGet).HandlerFunc(uh.getUserByID)
	usersR.Path("/{id}").Methods(http.MethodPut).HandlerFunc(uh.updateUser)
	usersR.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(uh.deleteUser)

	fmt.Println("Start listening")
	fmt.Println(http.ListenAndServe(":8080", util.LogMiddleware(r)))
}

func initDB() (*gorm.DB, error) {
	dataSourceName := "host=192.168.2.139 user=postgres password=1234 dbname=userdb port=5432"
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	db.AutoMigrate(&User{})

	return db, err
}

type userHandler struct {
	db *gorm.DB
}

func validateAndReturnFilterMap(filter string) (map[string]string, error) {
	splits := strings.Split(filter, ".")
	if len(splits) != 2 {
		return nil, errors.New("malformed sortBy query parameter, should be field.orderdirection")
	}

	field, value := splits[0], splits[1]

	if !stringInSlice(userFields, field) {
		return nil, errors.New("unknown field in filter query parameter")
	}

	return map[string]string{field: value}, nil
}

var userFields = getUserFields()

func getUserFields() []string {
	var field []string

	v := reflect.ValueOf(User{})
	for i := 0; i < v.Type().NumField(); i++ {
		field = append(field, v.Type().Field(i).Tag.Get("json"))
	}

	return field
}

func stringInSlice(strSlice []string, s string) bool {
	for _, v := range strSlice {
		if v == s {
			return true
		}
	}

	return false
}

func validateAndReturnSortQuery(sortBy string) (string, error) {
	splits := strings.Split(sortBy, ".")
	if len(splits) != 2 {
		return "", errors.New("malformed sortBy query parameter, should be field.orderdirection")
	}

	field, order := splits[0], splits[1]

	if order != "desc" && order != "asc" {
		return "", errors.New("malformed orderdirection in sortBy query parameter, should be asc or desc")
	}

	if !stringInSlice(userFields, field) {
		return "", errors.New("unknown field in sortBy query parameter")
	}

	return fmt.Sprintf("%s %s", field, strings.ToUpper(order)), nil

}

func (uh userHandler) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{}

	// sortBy is expected to look like field.orderdirection i. e. id.asc
	sortBy := r.URL.Query().Get("sortBy")
	if sortBy == "" {
		// id.asc is the default sort query
		sortBy = "id.asc"
	}

	sortQuery, err := validateAndReturnSortQuery(sortBy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	strLimit := r.URL.Query().Get("limit")
	// with a value as -1 for gorms Limit method, we'll get a request without limit as default
	limit := -1
	if strLimit != "" {
		limit, err = strconv.Atoi(strLimit)
		if err != nil || limit < -1 {
			http.Error(w, "limit query parameter is no valid number", http.StatusBadRequest)
			return
		}
	}

	strOffset := r.URL.Query().Get("offset")
	offset := -1
	if strOffset != "" {
		offset, err = strconv.Atoi(strOffset)
		if err != nil || offset < -1 {
			http.Error(w, "offset query parameter is no valid number", http.StatusBadRequest)
			return
		}
	}

	filter := r.URL.Query().Get("filter")
	filterMap := map[string]string{}
	if filter != "" {
		filterMap, err = validateAndReturnFilterMap(filter)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	if err := uh.db.Where(filterMap).Limit(limit).Offset(offset).Order(sortQuery).Find(&users).Error; err != nil {
		fmt.Println(err)
		http.Error(w, "Error on DB find for all users", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
	}
}

func (uh userHandler) getUserByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	user := User{ID: id}
	if err := uh.db.First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		fmt.Println(err)
		http.Error(w, fmt.Sprintf("Error on DB find for user with id: %s", id), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
	}
}

func (uh userHandler) updateUser(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	user := User{ID: id}
	if err := uh.db.First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		fmt.Println(err)
		http.Error(w, fmt.Sprintf("Error on DB find for user with id: %s", id), http.StatusInternalServerError)
		return
	}

	u := User{}
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		fmt.Println(err)
		http.Error(w, "Error deconding request body", http.StatusBadRequest)
		return
	}

	if err := uh.db.Save(&user).Error; err != nil {
		fmt.Println(err)
		http.Error(w, "Error saving user", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&u); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (uh userHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	user := User{ID: id}
	if err := uh.db.First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		fmt.Println(err)
		http.Error(w, fmt.Sprintf("Error on DB find for user with id: %s", id), http.StatusInternalServerError)
		return
	}

	if err := uh.db.Delete(&user).Error; err != nil {
		fmt.Println(err)
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (uh userHandler) createUser(w http.ResponseWriter, r *http.Request) {
	u := User{}

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := uh.db.Create(&u).Error; err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(&u)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

// TODO: change all error in initial version to http.Error instead of http.WriteHeader(error)
