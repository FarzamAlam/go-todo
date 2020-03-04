package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Title     string     `json:"title"`
	Priority  string     `gorm:"default:'0'" json:"priority"`
	Deadline  *time.Time `gorm:"default:null" json:"deadline"`
	Done      bool       `json:"done"`
	ProjectID uint       `json:"project_id"`
}

func (t *Task) Complete() {
	t.Done = true
}

func (t *Task) Undo() {
	t.Done = !t.Done
}

func GetAllTasks(project *Project) (*[]Task, error) {

	tasks := &[]Task{}
	err := GetDB().Model(project).Related(tasks).Error
	return tasks, err
}

func (task *Task) AddTask() error {
	err := GetDB().Save(task).Error
	if err != nil {
		return err
	}
	return nil
}
