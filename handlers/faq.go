package handlers

import (
	"fmt"
	"net/http"

	a "github.com/dariuszkorolczukcom/pinkbubbleapi/auth/v1"
	db "github.com/dariuszkorolczukcom/pinkbubbleapi/database"
	m "github.com/dariuszkorolczukcom/pinkbubbleapi/models"
	"github.com/gin-gonic/gin"
)

func GetFaq(c *gin.Context) {
	var faq []m.Faq
	db.Conn.Find(&faq)

	c.JSON(200, faq)
}

func AddFaq(c *gin.Context) {
	if !a.IsAdmin(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized to perform that operation"})
		return
	}
	var faq m.Faq

	if err := c.ShouldBindJSON(&faq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Conn.Create(&faq)
	db.Conn.Find(&faq)
	c.JSON(200, faq)
}

func UpdateFaq(c *gin.Context) {
	if !a.IsAdmin(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized to perform that operation"})
		return
	}

	var faq m.Faq

	if err := c.ShouldBindJSON(&faq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("%+v\n", faq)

	db.Conn.Save(&faq)
	db.Conn.First(&faq)

	c.JSON(200, faq)
}

func DeleteFaq(c *gin.Context) {
	if !a.IsAdmin(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized to perform that operation"})
		return
	}
	var faq m.Faq

	if err := c.ShouldBindJSON(&faq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Conn.Delete(&faq)
	db.Conn.First(&faq)

	c.JSON(200, faq)
}
