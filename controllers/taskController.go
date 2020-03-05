package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/farzamalam/go-todo/utils"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	"github.com/farzamalam/go-todo/models"
)

var CreateTask = func(w http.ResponseWriter, r *http.Request) {

	// Get title from the url and get project from using the title.
	title := mux.Vars(r)["title"]
	project, err, found := models.GetProject(title)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Error in GetProject")
		return
	}
	if !found {
		utils.RespondError(w, http.StatusBadRequest, "Error in the request")
		return
	}
	task := models.Task{ProjectID: project.ID}
	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Error in request body.")
		return
	}
	err = (&task).AddTask()
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Error while adding the task.")
		return
	}
	utils.Respond(w, http.StatusCreated, task)
}

var GetAllTask = func(w http.ResponseWriter, r *http.Request) {
	//To do for particular project.
	params := mux.Vars(r)
	projectTitle := params["title"]
	if projectTitle == "" {
		utils.RespondError(w, http.StatusBadRequest, "Title of the project is empty.")
		return
	}
	project, err, found := models.GetProject(projectTitle)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Error while getting project from title.")
		return
	}
	if !found {
		utils.RespondError(w, http.StatusBadRequest, "Project not found.")
		return
	}
	tasks, err := models.GetAllTasks(project)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Error in GetAllTasks.")
		return
	}
	utils.Respond(w, http.StatusOK, tasks)
}

var GetTaskID = func(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query()
	id, err := strconv.Atoi(queryParam.Get("id"))
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid id query param")
		return
	}
	task, err, found := models.GetTaskById(uint(id))
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Error in GetTaskById")
		return
	}
	if !found {
		utils.RespondError(w, http.StatusBadRequest, "Task is not found")
		return
	}
	utils.Respond(w, http.StatusOK, task)
}

var GetTaskTitle = func(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query()
	title := queryParam.Get("title")
	projectTitle := mux.Vars(r)["projectTitle"]
	if title == "" {
		utils.RespondError(w, http.StatusBadRequest, "Project Title is invalid")
		return
	}
	_, err, found := models.GetProject(projectTitle)

	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Error in GetProject")
		return
	}
	if !found {
		utils.RespondError(w, http.StatusBadRequest, "No project is found with the title")
		return
	}
	task, err, found := models.GetTask(title)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Error in GetTask")
		return
	}
	if !found {
		utils.RespondError(w, http.StatusBadRequest, "Task is not found")
		return
	}
	utils.Respond(w, http.StatusOK, task)
}

var DeleteTask = func(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query()
	id, err := strconv.Atoi(queryParam.Get("id"))
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Error in the query paramter")
		return
	}
	task := &models.Task{Model: gorm.Model{ID: uint(id)}}
	err = task.DeleteTask()
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Error in DeleteTask")
		return
	}
	utils.Respond(w, http.StatusAccepted, task)
}
