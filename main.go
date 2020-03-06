package main

import (
	"log"
	"net/http"
	"os"

	"github.com/farzamalam/go-todo/controllers"

	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
	"github.com/unrolled/secure"
)

func main() {
	router := mux.NewRouter()

	router.Use(secureMiddleware.Handler)

	err := godotenv.Load() // Load environment variable.
	if err != nil {
		log.Fatal("Error while getting the environment variable", err)
		os.Exit(2)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Println("Server started :", port)
	// Projects
	router.HandleFunc("/api/v1/projects", controllers.CreateProject).Methods("POST")
	router.HandleFunc("/api/v1/projects", controllers.GetAllProject).Methods("GET")
	router.HandleFunc("/api/v1/projects/{title}", controllers.GetProject).Methods("GET")
	router.HandleFunc("/api/v1/projects/{title}", controllers.UpdateProject).Methods("PUT")
	router.HandleFunc("/api/v1/projects/{title}", controllers.DeleteProject).Methods("DELETE")
	router.HandleFunc("/api/v1/projects/{title}", controllers.ArchieveProject).Methods("PATCH")
	// Tasks
	router.HandleFunc("/api/v1/projects/{title}/tasks", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/api/v1/projects/{title}/tasks", controllers.GetAllTask).Methods("GET")
	router.HandleFunc("/api/v1/projects/{projectTitle}/task", controllers.GetTaskID).Queries("id", "{id}").Methods("GET")
	router.HandleFunc("/api/v1/projects/{projectTitle}/task", controllers.DeleteTask).Queries("id", "{id}").Methods("DELETE")
	router.HandleFunc("/api/v1/projects/{projectTitle}/task", controllers.CreateTask).Queries("id", "{id}").Methods("PUT")
	router.HandleFunc("/api/v1/projects/{projectTitle}/task", controllers.GetTaskTitle).Queries("title", "{title}").Methods("GET")
	// Need to work.
	router.HandleFunc("/api/v1/projects/{title}/tasks", controllers.CreateTask).Methods("PUT")
	//router.Path("/api/v1/")
	log.Fatal(http.ListenAndServe(":"+port, router))

}

var secureMiddleware = secure.New(secure.Options{
	FrameDeny:      true,
	ReferrerPolicy: "same-origin",
})
