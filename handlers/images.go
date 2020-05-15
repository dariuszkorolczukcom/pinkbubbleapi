package handlers

import (
	"net/http"

	a "github.com/dariuszkorolczukcom/pinkbubbleapi/auth"
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
	// var buf bytes.Buffer
	// // in your case file would be fileupload
	// form, _ := c.MultipartForm()
	// files := form.File["file[]"]
	// file, header, err := files
	// // fmt.Println(file)
	// // fmt.Println(header)
	// check(err)
	// defer file.Close()

	// name := strings.Split(header.Filename, ".")
	// // io.Copy(&buf, file)
	// filepath := os.Getenv("FILEPATH")

	// err = ioutil.WriteFile(filepath+"/"+name[0]+".jpg", buf.Bytes(), 0644)
	// check(err)
	// fileName := name[0] + ".jpg"
	// // fmt.Printf("File name %s\n", name[0])
	// c.JSON(200, fileName)
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
