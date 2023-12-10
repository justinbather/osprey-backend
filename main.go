package main

import (
	"fmt"

	"github.com/gorilla/mux"

	"net/http"
	"osprey-backend/db"
	"osprey-backend/handlers"
)

func main() {

	if err := db.Connect(); err != nil {
		panic(err)
	}
	fmt.Println("MongoDB Connected")

	fmt.Println("Server running on port 3000")
	r := mux.NewRouter()

	r.HandleFunc("/log", handlers.NewLog).Methods("POST")
	r.HandleFunc("/projects", handlers.CreateProject).Methods("POST")
	r.HandleFunc("/projects/{id}", handlers.GetProject).Methods("GET")
	r.HandleFunc("/projects/{projId}/logs", handlers.GetLogs).Methods("GET")

	if err := http.ListenAndServe(":3000", r); err != nil {
		panic(err)
	}
}
