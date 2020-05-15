package handlers

import (
	"net/http"

	a "github.com/dariuszkorolczukcom/pinkbubbleapi/auth"
	db "github.com/dariuszkorolczukcom/pinkbubbleapi/database"
	m "github.com/dariuszkorolczukcom/pinkbubbleapi/models"
	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {

	var categories []m.Category
	db.Conn.Find(&categories)

	c.JSON(200, categories)
}

func GetCategory(c *gin.Context) {

	var category m.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Conn.First(&category)
}

func AddCategory(c *gin.Context) {
	if !a.IsAdmin(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized to perform that operation"})
		return
	}

	var category m.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Conn.Create(&category)
	db.Conn.Find(&category)
	c.JSON(200, category)
}

func UpdateCategory(c *gin.Context) {
	if !a.IsAdmin(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized to perform that operation"})
		return
	}

	var category m.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Conn.Save(&category)
	db.Conn.First(&category)

	c.JSON(200, category)
}

func DeleteCategory(c *gin.Context) {
	if !a.IsAdmin(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized to perform that operation"})
		return
	}

	var category m.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Conn.Delete(&category)
	db.Conn.First(&category)

	c.JSON(200, category)
}
