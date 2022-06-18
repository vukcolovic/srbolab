package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"srbolabApp/handlers"
	"time"
)

func runServer() {
	r := mux.NewRouter()

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("cao"))
	})

	s := r.PathPrefix("/api/users").Subrouter()
	s.HandleFunc("/register", handlers.Register).Methods("POST")
	s.HandleFunc("/list", handlers.ListUsers).Methods("GET")
	s.HandleFunc("/id/{id}", handlers.GetUserByID).Methods("GET")
	s.HandleFunc("/login", handlers.Login).Methods("POST")
	s.HandleFunc("/logout", handlers.Login).Methods("GET")
	log.Fatal(srv.ListenAndServe())
}
