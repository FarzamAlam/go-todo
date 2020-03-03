package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/farzamalam/go-todo/utils"

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
	tasks := models.GetAllTasks()
	utils.Respond(w, http.StatusOK, tasks)
}
