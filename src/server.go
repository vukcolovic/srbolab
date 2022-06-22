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
	r.Use(corsMiddleware)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("cao"))
	})

	//r.Use(middleware.AuthToken)

	s := r.PathPrefix("/api/users").Subrouter()
	s.HandleFunc("/register", handlers.Register).Methods("POST")
	s.HandleFunc("/list", handlers.ListUsers).Methods("GET")
	s.HandleFunc("/id/{id}", handlers.GetUserByID).Methods("GET")
	s.HandleFunc("/login", handlers.Login).Methods("POST")
	s.HandleFunc("/logout", handlers.Login).Methods("GET")
	log.Fatal(srv.ListenAndServe())
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8080")
		//w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Add("Access-Control-Allow-Headers", "X-Requested-With, X-HTTP-Method-Override, Content-Type, Accept")
		//w.Header().Set("content-type", "application/json;charset=UTF-8")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
