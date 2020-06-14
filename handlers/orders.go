package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	a "github.com/dariuszkorolczukcom/pinkbubbleapi/auth"
	db "github.com/dariuszkorolczukcom/pinkbubbleapi/database"
	m "github.com/dariuszkorolczukcom/pinkbubbleapi/models"
	"github.com/gin-gonic/gin"
)

func GetOrder(c *gin.Context) {
	id := a.GetUserId(c)
	orderid, err := strconv.Atoi(c.Param("orderid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var order m.Order
	var user m.User

	db.Conn.First(&user, id)
	db.Conn.First(&order, orderid)
	db.Conn.Model(&order).Related(&order.OrderItems)
	for j, _ := range order.OrderItems {
		q := &order.OrderItems[j]
		db.Conn.Model(&q).Related(&q.Product)
	}
	c.JSON(200, order)
}

func GetOrders(c *gin.Context) {
	id := a.GetUserId(c)
	var orders []m.Order
	var user m.User

	db.Conn.Where("user_id=?", id).Find(&orders)
	db.Conn.Where("user_id=?", id).Find(&user)
	for i, _ := range orders {
		p := &orders[i]
		db.Conn.Model(&p).Related(&p.OrderItems)
		for j, _ := range p.OrderItems {
			q := &p.OrderItems[j]
			db.Conn.Model(&q).Related(&q.Product)
		}
	}

	c.JSON(200, orders)
}

func AddOrder(c *gin.Context) {
	id := a.GetUserId(c)
	var order m.Order
	var user m.User

	db.Conn.First(&user, id)

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order.User = user
	order.Status = 1

	for j, _ := range order.OrderItems {
		q := &order.OrderItems[j]
		db.Conn.Model(&q).Related(&q.Product)
	}
	fmt.Printf("%+v\n", order)

	db.Conn.Create(&order)
	db.Conn.Find(&order)

	c.JSON(200, order)
}

func AddGuestOrder(c *gin.Context) {
	var order m.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for j, _ := range order.OrderItems {
		q := &order.OrderItems[j]
		db.Conn.Model(&q).Related(&q.Product)
	}
	fmt.Printf("%+v\n", order)

	db.Conn.Save(&order)

	c.JSON(200, order)
}
