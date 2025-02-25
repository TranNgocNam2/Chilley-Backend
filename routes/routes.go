package routes

import (
	"chilley.nam2507/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/tasks", handlers.GetTasks())
	r.POST("/tasks", handlers.AddTask())
	r.PUT("/tasks/:id", handlers.UpdateTask())
	r.DELETE("/tasks/:id", handlers.DeleteTask())
}
