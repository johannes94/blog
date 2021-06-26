package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/user"

	"github.com/gorilla/mux"
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

func main() {

	db, err := initDB()
	if err != nil {
		fmt.Printf("Error initializing DB: %v\n", err)
	}

	uh := userHandler{db: db}

	r := mux.NewRouter()
	usersR := r.PathPrefix("/users").Subrouter()
	usersR.Path("").Methods(http.MethodGet).HandlerFunc(uh.getAllUsers)
	usersR.Path("").Methods(http.MethodPost).HandlerFunc(uh.createUser)
	usersR.Path("/{id}").Methods(http.MethodGet).HandlerFunc(uh.getUserByID)
	usersR.Path("/{id}").Methods(http.MethodPut).HandlerFunc(uh.updateUser)
	usersR.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(uh.deleteUser)

	fmt.Println("Start listening")
	fmt.Println(http.ListenAndServe(":8080", r))
}

func initDB() (*gorm.DB, error) {
	dataSourceName := "host=192.168.2.139 user=postgres password=1234 dbname=userdb port=5432"
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	db.AutoMigrate(&user.User{})

	return db, err
}

type userHandler struct {
	db *gorm.DB
}

func (uh userHandler) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{}
	if err := uh.db.Find(&users).Error; err != nil {
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
