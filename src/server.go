package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"srbolabApp/handlers"
	"srbolabApp/middleware"
	"time"
)

func runServer() {
	r := mux.NewRouter()
	r.Use(middleware.CorsMiddleware)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("cao"))
	})

	r.Use(middleware.AuthToken)

	s := r.PathPrefix("/api/users").Subrouter()
	s.HandleFunc("/register", handlers.Register).Methods("POST")
	s.HandleFunc("/list", handlers.ListUsers).Methods("GET")
	s.HandleFunc("/id/{id}", handlers.GetUserByID).Methods("GET")
	s.HandleFunc("/login", handlers.Login).Methods("POST")
	s.HandleFunc("/delete/{id}", handlers.DeleteUser).Methods("GET")
	s.HandleFunc("/count", handlers.CountUsers).Methods("GET")

	log.Fatal(srv.ListenAndServe())
}
