package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/johannes94/posts/crud-rest-api-with-gorilla-mux/initial/user"
)

func main() {
	r := mux.NewRouter()
	usersR := r.PathPrefix("/users").Subrouter()
	usersR.Path("").Methods(http.MethodGet).HandlerFunc(user.GetAllUsers)
	usersR.Path("").Methods(http.MethodPost).HandlerFunc(user.CreateUser)
	usersR.Path("/{id}").Methods(http.MethodGet).HandlerFunc(user.GetUserByID)
	usersR.Path("/{id}").Methods(http.MethodPut).HandlerFunc(user.UptdateUser)
	usersR.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(user.DeleteUser)

	fmt.Println("Start listening")
	fmt.Println(http.ListenAndServe(":8080", r))
}
