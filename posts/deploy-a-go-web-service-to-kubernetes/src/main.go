package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", httpHandleFunc)
	http.ListenAndServe(":8080", nil)
}

func httpHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("was called")
	w.Write([]byte("<h1>Hello World</h1>"))
}
