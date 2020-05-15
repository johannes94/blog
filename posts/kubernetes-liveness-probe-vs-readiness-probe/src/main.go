package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	time.Sleep(time.Second * time.Duration(30))
	ioutil.WriteFile("init.txt", []byte("Initialization done."), 0644)

	healthy := true

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Called at: %v\n", time.Now())
		if healthy {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("healthy"))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("unhealthy"))
		}
	})

	http.HandleFunc("/kill", func(w http.ResponseWriter, r *http.Request) {
		healthy = false
		w.Write([]byte("Set to unhealthy"))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
