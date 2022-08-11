package main

import (
	"net/http"
	"ourtask/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	_ = r.Run()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	userRepo := controllers.New()
	r.POST("/tasks", userRepo.CreateTask)
	r.GET("/tasks", userRepo.GetTasks)
	r.GET("/tasks/:id", userRepo.GetTask)
	r.PUT("/tasks/:id", userRepo.UpdateTask)
	r.DELETE("/tasks/:id", userRepo.DeleteTask)

	return r
}
