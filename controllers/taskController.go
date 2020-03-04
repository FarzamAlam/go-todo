package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/farzamalam/go-todo/utils"
	"github.com/gorilla/mux"

	"github.com/farzamalam/go-todo/models"
)

var CreateTask = func(w http.ResponseWriter, r *http.Request) {
	task := &models.Task{}
	err := json.NewDecoder(r.Body).Decode(task)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Error in request body.")
	}
	err = task.AddTask()
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Error while adding the task.")
	}
	utils.Respond(w, http.StatusCreated, task)
}

var GetAllTask = func(w http.ResponseWriter, r *http.Request) {
	//To do for particular project.
	params := mux.Vars(r)
	projectTitle := params["title"]
	if projectTitle == "" {
		utils.RespondError(w, http.StatusBadRequest, "Title of the project is empty.")
	}
	project, err, found := models.GetProject(projectTitle)
	log.Println("found : ", found)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Error while getting project from title.")
		return
	}
	if !found {
		log.Println("Inside project not found")
		utils.RespondError(w, http.StatusBadRequest, "Project not found.")
		return
	}
	log.Println("After the RespondError")
	tasks, err := models.GetAllTasks(project)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Error in GetAllTasks.")
		return
	}
	utils.Respond(w, http.StatusOK, tasks)
}
