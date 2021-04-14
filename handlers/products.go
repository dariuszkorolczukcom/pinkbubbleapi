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

func GetSortedProducts(c *gin.Context) {

	var categories []m.Category
	db.Conn.Find(&categories)

	for i, _ := range categories {
		ca := &categories[i]
		db.Conn.Model(&ca).Related(&ca.Products)
		for i, _ := range ca.Products {
			p := &ca.Products[i]
			db.Conn.Model(&p).Related(&p.Images)
		}
	}

	c.JSON(200, categories)
}

func GetProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var product m.Product

	db.Conn.First(&product, id)
	db.Conn.Model(&product).Related(&product.Images)
	c.JSON(200, product)
}

func GetProducts(c *gin.Context) {

	var products []m.Product
	db.Conn.Find(&products)

	for i, _ := range products {
		p := &products[i]
		db.Conn.Model(&p).Related(&p.Images)
	}

	c.JSON(200, products)
}

func AddProduct(c *gin.Context) {
	if !a.IsAdmin(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized to perform that operation"})
		return
	}
	var product m.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("%+v\n", product)

	db.Conn.Create(&product)
	db.Conn.Find(&product)
	c.JSON(200, product)
}

func UpdateProduct(c *gin.Context) {
	if !a.IsAdmin(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized to perform that operation"})
		return
	}

	var product m.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("%+v\n", product)

	db.Conn.Save(&product)
	for _, img := range product.Images {
		db.Conn.FirstOrCreate(&img)
	}
	db.Conn.First(&product)

	c.JSON(200, product)
}

func DeleteProduct(c *gin.Context) {
	if !a.IsAdmin(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized to perform that operation"})
		return
	}
	var product m.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Conn.Delete(&product)
	db.Conn.First(&product)
	fmt.Println(product)

	c.JSON(200, product)
}
