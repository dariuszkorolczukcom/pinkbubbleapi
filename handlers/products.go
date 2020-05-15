package handlers

import (
	"fmt"
	"net/http"

	a "github.com/dariuszkorolczukcom/pinkbubbleapi/auth"
	db "github.com/dariuszkorolczukcom/pinkbubbleapi/database"
	m "github.com/dariuszkorolczukcom/pinkbubbleapi/models"
	"github.com/gin-gonic/gin"
)

func GetSortedProducts(c *gin.Context) {

	var categories []m.Category
	db.Conn.Find(&categories)

	for i, _ := range categories {
		c := &categories[i]
		db.Conn.Model(&c).Related(&c.Products)
		for i, _ := range c.Products {
			p := &c.Products[i]
			db.Conn.Model(&p).Related(&p.Images)
		}
	}

	c.JSON(200, categories)
}

func GetProduct(c *gin.Context) {
	var product m.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Conn.First(&product)

	c.JSON(200, product)
}

func GetProducts(c *gin.Context) {

	var products []m.Product
	db.Conn.Find(&products)

	for _, p := range products {
		var images []m.Image
		db.Conn.Model(&p).Related(&images)
		p.Images = images
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

// [
//     {
//         "name": "bath",
//         "items": [
//             {
//                 "id": 1,
//                 "img": "bomb.jpg",
//                 "price": 4.99,
//                 "title": "card title",
//                 "subtitle": "card subtitle",
//                 "text":
//                     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam at massa ac leo congue tincidunt vel at velit."
//             },
//             {
//                 "id": 2,
//                 "category": 1,
//                 "img": "bomb.jpg",
//                 "price": 4.99,
//                 "title": "card title 2",
//                 "subtitle": "card subtitle",
//                 "text":
//                     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam at massa ac leo congue tincidunt vel at velit."
//             },
//             {
//                 "id": 3,
//                 "category": 1,
//                 "img": "bomb.jpg",
//                 "price": 4.99,
//                 "title": "card title 3",
//                 "subtitle": "card subtitle",
//                 "text":
//                     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam at massa ac leo congue tincidunt vel at velit."
//             },
//             {
//                 "id": 3,
//                 "category": 1,
//                 "title": "card title 3",
//                 "price": 4.99,
//                 "subtitle": "card subtitle",
//                 "text":
//                     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam at massa ac leo congue tincidunt vel at velit."
//             }
//         ]
//     },
//     {
//         "name": "beauty",
//         "items": [
//             {
//                 "id": 3,
//                 "category": 2,
//                 "title": "card title 3",
//                 "price": 4.99,
//                 "subtitle": "card subtitle",
//                 "text":
//                     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam at massa ac leo congue tincidunt vel at velit."
//             },
//             {
//                 "id": 3,
//                 "category": 2,
//                 "title": "card title 3",
//                 "subtitle": "card subtitle",
//                 "text":
//                     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam at massa ac leo congue tincidunt vel at velit."
//             },
//             {
//                 "id": 3,
//                 "category": 2,
//                 "title": "card title 3",
//                 "subtitle": "card subtitle",
//                 "text":
//                     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam at massa ac leo congue tincidunt vel at velit."
//             },
//             {
//                 "id": 3,
//                 "category": 2,
//                 "title": "card title 3",
//                 "subtitle": "card subtitle",
//                 "text":
//                     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam at massa ac leo congue tincidunt vel at velit."
//             },
//             {
//                 "id": 3,
//                 "category": 2,
//                 "title": "card title 3",
//                 "subtitle": "card subtitle",
//                 "text":
//                     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam at massa ac leo congue tincidunt vel at velit."
//             },
//             {
//                 "id": 3,
//                 "category": 2,
//                 "title": "card title 3",
//                 "subtitle": "card subtitle",
//                 "text":
//                     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam at massa ac leo congue tincidunt vel at velit."
//             },
//             {
//                 "id": 3,
//                 "category": 2,
//                 "title": "card title 3",
//                 "subtitle": "card subtitle",
//                 "text":
//                     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam at massa ac leo congue tincidunt vel at velit."
//             }
//         ]
//     },
//     {
//         "name": "perfumes",
//         "items": [
//             {
//                 "id": 3,
//                 "category": 3,
//                 "title": "card title 3",
//                 "subtitle": "card subtitle",
//                 "text":
//                     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam at massa ac leo congue tincidunt vel at velit."
//             }
//         ]
//     }
// ]
