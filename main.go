package main

import (
	"log"
	"net/http"
	"os"

	"github.com/farzamalam/go-todo/controllers"

	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load() // Load environment variable.
	if err != nil {
		log.Fatal("Error while getting the environment variable", err)
		os.Exit(2)
	}
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/projects", controllers.CreateProject).Methods("POST")
	router.HandleFunc("/api/v1/projects", controllers.GetAllProject).Methods("GET")

	router.HandleFunc("/api/v1/projects/{title}/tasks", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/api/v1/projects/{title}/tasks", controllers.GetAllTask).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))

}
