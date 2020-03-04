package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/farzamalam/go-todo/utils"

	"github.com/farzamalam/go-todo/models"
)

var CreateProject = func(w http.ResponseWriter, r *http.Request) {
	project := &models.Project{}

	err := json.NewDecoder(r.Body).Decode(project)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Error in the request body.")
		return
	}
	err = project.AddProject()
	if err != nil {
		log.Println("Error in CreateProject : ", err)
		utils.RespondError(w, http.StatusInternalServerError, "Error while adding project")
		return
	}
	utils.Respond(w, http.StatusOK, project)
}

var GetAllProject = func(w http.ResponseWriter, r *http.Request) {
	projects := models.GetAllProjects()
	utils.Respond(w, http.StatusOK, projects)
}
