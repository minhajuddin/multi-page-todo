package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type requestHandler struct{}

func (h *requestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s\t[%s]\t%s\n", r.RemoteAddr, r.Method, r.URL.Path)
	fmt.Fprintf(w, "<!doctype html><h1>Welcome to the thunderdome!</h1>")
}

func main() {
	log.Println("Starting Multi Page TODO")

	server := &http.Server{
		Addr:         ":8080",
		Handler:      &requestHandler{},
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
