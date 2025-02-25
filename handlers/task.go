package handlers

import (
	"chilley.nam2507/models"
	"chilley.nam2507/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetTasks() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, services.GetTask())
	}
}

func AddTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var task models.Task
		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input!"})
			return
		}

		if task.Title == "" || task.Description == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Title and description are required!"})
			return
		}

		id := services.AddTask(task)
		c.JSON(200, gin.H{"id": id})
	}
}

func UpdateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task id!"})
			return
		}
		var req struct {
			Completed bool `json:"completed"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input!"})
			return
		}

		if !services.UpdateTask(id, req.Completed) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found!"})
			return
		}
		c.JSON(200, gin.H{"message": "Task updated successfully!"})
	}
}

func DeleteTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task id!"})
			return
		}
		if !services.DeleteTask(id) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found!"})
			return
		}
		c.JSON(200, gin.H{"message": "Task deleted successfully!"})
	}
}
