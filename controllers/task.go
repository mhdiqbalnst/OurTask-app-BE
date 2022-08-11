package controllers

import (
	"errors"
	"net/http"
	"ourtask/models"
	"strconv"

	"ourtask/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// task repository structs
type TaskRepo struct {
	Db *gorm.DB
}

func New() *TaskRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Task{})
	return &TaskRepo{Db: db}
}

// create task
func (repository *TaskRepo) CreateTask(c *gin.Context) {
	var task models.Task
	c.BindJSON(&task)
	err := models.CreateTask(repository.Db, &task)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, task)
}

// get tasks
func (repository *TaskRepo) GetTasks(c *gin.Context) {
	var task []models.Task
	err := models.GetTasks(repository.Db, &task)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, task)
}

// get task by id
func (repository *TaskRepo) GetTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var task models.Task
	err := models.GetTask(repository.Db, &task, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, task)
}

// update task
func (repository *TaskRepo) UpdateTask(c *gin.Context) {
	var task models.Task
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.GetTask(repository.Db, &task, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&task)
	err = models.UpdateTask(repository.Db, &task)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, task)
}

// delete task

func (repository *TaskRepo) DeleteTask(c *gin.Context) {
	var task models.Task
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteTask(repository.Db, &task, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
