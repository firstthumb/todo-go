package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/firstthumb/todo/models"
	"github.com/gin-gonic/gin"
)

func GetTasks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, models.GetTasks(db))
	}
}

func PutTask(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var task models.Task
		c.BindJSON(&task)
		id, err := models.PutTask(db, task.Name)
		if err == nil {
			c.JSON(http.StatusCreated, gin.H{
				"created": id,
			})
		} else {
			c.JSON(http.StatusNotModified, gin.H{
				"error": "Could not create task",
			})
		}
	}
}

func DeleteTask(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		_, err := models.DeleteTask(db, id)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"deleted": id,
			})
		} else {
			c.JSON(http.StatusNotModified, gin.H{
				"error": "Could not delete task",
			})
		}
	}
}
