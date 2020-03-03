package models

import (
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
