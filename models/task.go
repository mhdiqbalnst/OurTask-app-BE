package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID          int
	Detail_task string
	Member      string
	Deadline    string
	Progress    string
}

// create a task
func CreateTask(db *gorm.DB, Task *Task) (err error) {
	err = db.Create(Task).Error
	if err != nil {
		return err
	}
	return nil
}

// get tasks
func GetTasks(db *gorm.DB, Task *[]Task) (err error) {
	err = db.Find(Task).Error
	if err != nil {
		return err
	}
	return nil
}

// get task by id
func GetTask(db *gorm.DB, Task *Task, id int) (err error) {
	err = db.Where("id = ?", id).First(Task).Error
	if err != nil {
		return err
	}
	return nil
}

// update task
func UpdateTask(db *gorm.DB, Task *Task) (err error) {
	db.Save(Task)
	return nil
}

// delete task
func DeleteTask(db *gorm.DB, Task *Task, id int) (err error) {
	db.Where("id = ?", id).Delete(Task)
	return nil
}
