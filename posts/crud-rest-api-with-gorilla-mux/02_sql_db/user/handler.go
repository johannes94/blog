package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var users = []User{}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(users); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	index := indexByID(users, id)

	if index < 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users[index]); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func UptdateUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	index := indexByID(users, id)
	if index < 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	u := User{}
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	users[index] = u

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&u); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	index := indexByID(users, id)
	if index < 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	users = append(users[:index], users[index+1:]...)
	w.WriteHeader(http.StatusOK)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	u := User{}

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	users = append(users, u)

	response, err := json.Marshal(&u)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func indexByID(users []User, id string) int {
	for i := 0; i < len(users); i++ {
		if users[i].ID == id {
			return i
		}
	}

	return -1
}
