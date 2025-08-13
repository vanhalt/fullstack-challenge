package main

import (
	"log"
	"net/http"
	"task-api/handlers"

	"github.com/gorilla/mux"
)


func main() {
	r := mux.NewRouter()

	// Endpoints
	r.HandleFunc("/tasks",  handlers.ListTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", handlers.GetTask).Methods("GET")
	r.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")
	cors := handlers.CreateCORSHandler()

	log.Println("Server starting on port 3001...")
	log.Fatal(http.ListenAndServe(":3001", cors(r)))
}
