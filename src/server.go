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
		w.Write([]byte(""))
	})

	r.Use(middleware.AuthToken)

	s := r.PathPrefix("/api/users").Subrouter()
	s.HandleFunc("/register", handlers.Register).Methods("POST")
	s.HandleFunc("/update", handlers.UpdateUser).Methods("POST")
	s.HandleFunc("/list", handlers.ListUsers).Methods("GET")
	s.HandleFunc("/id/{id}", handlers.GetUserByID).Methods("GET")
	s.HandleFunc("/login", handlers.Login).Methods("POST")
	s.HandleFunc("/delete/{id}", handlers.DeleteUser).Methods("GET")
	s.HandleFunc("/count", handlers.CountUsers).Methods("GET")

	s = r.PathPrefix("/api/enumeration").Subrouter()
	s.HandleFunc("/irregularity-levels/all", handlers.ListIrregularityLevels).Methods("GET")

	s = r.PathPrefix("/api/irregularity").Subrouter()
	s.HandleFunc("/id/{id}", handlers.GetIrregularityByID).Methods("GET")
	s.HandleFunc("/create", handlers.CreateIrregularity).Methods("POST")
	s.HandleFunc("/list", handlers.ListIrregularities).Methods("POST")
	s.HandleFunc("/delete/{id}", handlers.DeleteIrregularity).Methods("GET")
	s.HandleFunc("/count", handlers.CountIrregularities).Methods("POST")
	s.HandleFunc("/update", handlers.UpdateIrregularities).Methods("POST")
	s.HandleFunc("/change-corrected", handlers.ChangeCorrected).Methods("POST")

	s = r.PathPrefix("/api/fuel-consumption").Subrouter()
	s.HandleFunc("/id/{id}", handlers.GetFuelConsumptionByID).Methods("GET")
	s.HandleFunc("/create", handlers.CreateFuelConsumption).Methods("POST")
	s.HandleFunc("/list", handlers.ListFuelConsumptions).Methods("POST")
	s.HandleFunc("/delete/{id}", handlers.DeleteFuelConsumption).Methods("GET")
	s.HandleFunc("/count", handlers.CountFuelConsumptions).Methods("POST")
	s.HandleFunc("/update", handlers.UpdateFuelConsumption).Methods("POST")
	s.HandleFunc("/sum-price", handlers.CountSumPrice).Methods("POST")

	s = r.PathPrefix("/api/certificate").Subrouter()
	s.HandleFunc("/id/{id}", handlers.GetCertificateByID).Methods("GET")
	s.HandleFunc("/create", handlers.CreateCertificate).Methods("POST")
	s.HandleFunc("/list", handlers.ListCertificates).Methods("POST")
	s.HandleFunc("/delete/{id}", handlers.DeleteCertificate).Methods("GET")
	s.HandleFunc("/count", handlers.CountCertificates).Methods("POST")
	s.HandleFunc("/update", handlers.UpdateCertificate).Methods("POST")
	s.HandleFunc("/pdf/id/{id}", handlers.GetPdfReportById).Methods("GET")

	log.Fatal(srv.ListenAndServe())
}
