package handlers

import (
	"net/http"
	"os"

	a "github.com/dariuszkorolczukcom/pinkbubbleapi/auth/v1"
	aws "github.com/dariuszkorolczukcom/pinkbubbleapi/aws"
	db "github.com/dariuszkorolczukcom/pinkbubbleapi/database"
	m "github.com/dariuszkorolczukcom/pinkbubbleapi/models"
	"github.com/gin-gonic/gin"
)

func GetImages(c *gin.Context) {

	var images []m.Image
	db.Conn.Find(&images)

	c.JSON(200, images)
}

func GetImage(c *gin.Context) {
	var image m.Image
	if err := c.ShouldBindJSON(&image); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Conn.First(&image)
}

func AddImage(c *gin.Context) {
	if !a.IsAdmin(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized to perform that operation"})
		return
	}

	file, header, err := c.Request.FormFile("file")
	check(err)
	defer file.Close()

	fileName := header.Filename

	if err = aws.UploadToS3(file, fileName); err != nil {
		c.JSON(500, gin.H{"error uploading file to s3": err.Error()})
		return
	}

	c.JSON(200, fileName)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func AddImageToProduct(c *gin.Context) {
	if !a.IsAdmin(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized to perform that operation"})
		return
	}

	var image m.Image
	if err := c.ShouldBindJSON(&image); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Conn.Create(&image)
	db.Conn.Find(&image)
	c.JSON(200, image)
}

func DeleteImage(c *gin.Context) {
	if !a.IsAdmin(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized to perform that operation"})
		return
	}

	var image m.Image
	if err := c.ShouldBindJSON(&image); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Conn.Delete(&image)
	db.Conn.First(&image)

	c.JSON(200, image)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
