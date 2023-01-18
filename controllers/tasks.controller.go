package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/roalcantara/api/db"
	"github.com/roalcantara/api/models"
)

func TaskIndex(c *gin.Context) {
	var tasks []models.Task
	result := db.DB.Find(&tasks)

	if result.Error != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func TaskPending(c *gin.Context) {
	var tasks []models.Task
	result := db.DB.Where("done = ?", false).Find(&tasks)

	if result.Error != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func TaskCreate(c *gin.Context) {
	var body struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Done        bool   `json:"done"`
	}
	c.Bind(&body)

	task := models.Task{Title: body.Title, Description: body.Description}
	result := db.DB.Create(&task)

	if result.Error != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	c.JSON(http.StatusCreated, task)
}

func TaskUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Done        bool   `json:"done"`
	}
	c.Bind(&body)

	var task models.Task
	db.DB.First(&task, id)
	result := db.DB.Model(&task).Updates(models.Task{
		Title:       body.Title,
		Description: body.Description,
		Done:        body.Done,
	})

	if result.Error != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	c.JSON(http.StatusAccepted, task)
}

func TaskShow(c *gin.Context) {
	id := c.Param("id")

	var task models.Task
	result := db.DB.Find(&task, id)

	if result.Error != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	c.JSON(http.StatusFound, task)
}

func TaskDelete(c *gin.Context) {
	id := c.Param("id")

	var task models.Task
	result := db.DB.Delete(&task, id)

	if result.Error != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	c.JSON(http.StatusAccepted, task)
}

func AddTaskRoutes(r *gin.Engine) {
	r.GET("/api/tasks", TaskPending)
	r.GET("/api/tasks/all", TaskIndex)
	r.POST("/api/tasks", TaskCreate)
	r.GET("/api/tasks/:id", TaskShow)
	r.PUT("/api/tasks/:id", TaskUpdate)
	r.DELETE("/api/tasks/:id", TaskDelete)
}
