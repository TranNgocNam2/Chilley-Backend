package services

import (
	"chilley.nam2507/models"
)

var tasks []models.Task
var lastTaskID = 0

func GetTask() []models.Task {
	return tasks
}

func AddTask(newTask models.Task) int {
	lastTaskID++
	newTask.ID = lastTaskID
	tasks = append(tasks, newTask)
	return newTask.ID
}

func UpdateTask(id int, completed bool) bool {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = completed
			return true
		}
	}
	return false
}

func DeleteTask(id int) bool {
	for i, task := range tasks {
		if task.ID == id {
			if len(tasks) == 1 {
				tasks = nil
				return true
			}
			tasks = append(tasks[:i], tasks[i+1])
			return true
		}
	}
	return false
}
