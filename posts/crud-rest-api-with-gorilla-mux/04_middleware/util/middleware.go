package util

import (
	"fmt"
	"net/http"
)

func ContentTypeMiddleware(contentType string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(rw, r)
		rw.Header().Add("Content-Type", contentType)
	})
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Printf("> Request from %v - Method: %v - Path: %v\n", r.RemoteAddr, r.Method, r.URL.Path)
		next.ServeHTTP(rw, r)
	})
}
