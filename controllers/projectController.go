package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/farzamalam/go-todo/utils"
	"github.com/gorilla/mux"

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

var GetProject = func(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	if title == "" {
		utils.RespondError(w, http.StatusBadRequest, "Project Title is invalid")
		return
	}
	project, err, found := models.GetProject(title)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Error in GetProject")
		return
	}
	if !found {
		utils.RespondError(w, http.StatusBadRequest, "No project is found with the title")
		return
	}
	utils.Respond(w, http.StatusOK, project)
}

var UpdateProject = func(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	if title == "" {
		utils.RespondError(w, http.StatusBadRequest, "Project title is invalid")
		return
	}
	project, err, found := models.GetProject(title)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Error while getting the project")
		return
	}
	if !found {
		utils.RespondError(w, http.StatusBadRequest, "Project not found with title.")
		return
	}
	err = json.NewDecoder(r.Body).Decode(project)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Error while marhsalling the request body.")
		return
	}
	err = project.UpdateProject()
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Error while updating the project")
		return
	}
	utils.Respond(w, http.StatusAccepted, project)
}

var DeleteProject = func(w http.ResponseWriter, r *http.Request) {

}
