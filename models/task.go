package models

import (
	"log"
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

func GetAllTasks() *[]Task {
	tasks := &[]Task{}
	GetDB().Find(tasks)
	return tasks
}

func (task *Task) AddTask() error {
	err := GetDB().Save(task).Error
	if err != nil {
		log.Fatal("Error while adding task : ", err)
		return err
	}
	return nil
}
