package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Project struct {
	gorm.Model
	Title     string `gorm:"unique" json:"title"`
	Archieved bool   `json:"archieved"`
	Tasks     []Task `gorm:"ForeignKey:ProjectID" json:"tasks"`
}

// GetAllProjects ... It will fetch all the projects from the table and return them.
func GetAllProjects() *[]Project {
	projects := &[]Project{}
	GetDB().Find(projects)
	return projects
}

func (project *Project) AddProject() error {
	err := GetDB().Save(project).Error
	if err != nil {
		return err
	}
	return nil
}

func GetProject(title string) (*Project, error, bool) {
	project := &Project{}
	err := GetDB().First(project, Project{Title: title}).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil, false
		} else {
			return nil, err, false
		}

	}
	return project, nil, true
}

func (project *Project) DeleteProject() error {
	err := GetDB().Delete(project).Error
	if err != nil {
		log.Println("Error in DeleteProject : ", err)
		return err
	}
	return nil
}

func (project *Project) ArchieveProject() {
	project.Archieved = !project.Archieved
}
