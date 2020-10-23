package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/johannes94/posts/crud-rest-api-with-gorilla-mux/initial/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	r := mux.NewRouter()

	db, err := initDB()
	if err != nil {
		log.Fatalf("initDB failed with error: %v", err)
	}

	usersR := r.PathPrefix("/users").Subrouter()
	uh := user.Handler{DB: db}
	usersR.Path("").Methods(http.MethodGet).HandlerFunc(uh.GetAllUsers)
	usersR.Path("").Methods(http.MethodPost).HandlerFunc(uh.CreateUser)
	usersR.Path("/{id}").Methods(http.MethodGet).HandlerFunc(uh.GetUserByID)
	usersR.Path("/{id}").Methods(http.MethodPut).HandlerFunc(uh.UpdateUser)
	usersR.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(uh.DeleteUser)

	fmt.Println("Start listening")
	fmt.Println(http.ListenAndServe(":8080", r))
}

func initDB() (*gorm.DB, error) {
	dataSourceName := "host=localhost user=postgres password=1234 dbname=userdb port=5432"
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	db.AutoMigrate(&user.User{})

	return db, err
}
