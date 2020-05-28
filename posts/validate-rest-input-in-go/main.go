package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type User struct {
	ID                 int    `validate:"isdefault"`
	FirstName          string `validate:"required"`
	LastName           string `validate:"required"`
	FavouriteVideoGame string `validate:"game-blacklist"`
	Email              string `validate:"required,email"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user", PostUser).Methods(http.MethodPost)
	router.HandleFunc("/user", GetUsers).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8081", router))
}

var users = []User{}
var id = 0

func GameBlacklistValidator(f1 validator.FieldLevel) bool {
	gameBlacklist := []string{"PUBG", "Fortnite"}
	game := f1.Field().String()
	for _, g := range gameBlacklist {
		if game == g {
			return false
		}
	}
	return true
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	user := User{}
	json.NewDecoder(r.Body).Decode(&user)

	validate := validator.New()
	validate.RegisterValidation("game-blacklist", GameBlacklistValidator)

	if err := validate.Struct(user); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		responseBody := map[string]string{"error": validationErrors.Error()}
		if err := json.NewEncoder(w).Encode(responseBody); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	// We don't want an API user to set the ID manually
	// in a production use case this could be an automatically
	// ID in the database
	user.ID = id
	id++

	users = append(users, user)
	w.WriteHeader(http.StatusCreated)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
