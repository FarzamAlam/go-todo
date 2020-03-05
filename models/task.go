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

func (task *Task) DeleteTask() error {
	err := GetDB().Delete(task).Error
	if err != nil {
		return err
	}
	return nil
}
func GetTask(title string) (*Task, error, bool) {
	task := &Task{}
	err := GetDB().First(task, Task{Title: title}).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil, false
		}
		return nil, err, false

	}
	return task, nil, true
}
func GetTaskById(id uint) (*Task, error, bool) {
	task := &Task{}
	err := GetDB().First(task, Task{ProjectID: id}).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil, false
		}
		return nil, err, false
	}
	return task, nil, true
}
