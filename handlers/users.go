package handlers

import (
	db "github.com/dariuszkorolczukcom/pinkbubbleapi/database"
	m "github.com/dariuszkorolczukcom/pinkbubbleapi/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

	var users []m.User
	db.Conn.Find(&users)

	c.JSON(200, users)
}
