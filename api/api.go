package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	a "github.com/dariuszkorolczukcom/pinkbubbleapi/auth"
	db "github.com/dariuszkorolczukcom/pinkbubbleapi/database"
	h "github.com/dariuszkorolczukcom/pinkbubbleapi/handlers"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func homeLink(c *gin.Context) {
	c.JSON(http.StatusOK, "I am an API!\nBuild by an intelligent being,\nRunning on a stupid machine.\nv1.0.0")
}

func main() {
	db.Conn = db.InitDB()

	defer db.Conn.Close()
	r := gin.Default()

	r.Use(cors.New(CORS()))
	r.Use(Logger())
	r.Use(gin.Recovery())
	authMiddleware, err := a.AuthMiddleware()
	authMiddlewareClient, err := a.AuthMiddlewareClient()
	auth := r.Group("/auth")
	admin := r.Group("/admin")
	order := r.Group("/order")
	info := r.Group("/info")
	info.Use(authMiddlewareClient.MiddlewareFunc())
	auth.Use(authMiddleware.MiddlewareFunc())
	admin.Use(authMiddleware.MiddlewareFunc())
	order.Use(authMiddlewareClient.MiddlewareFunc())
	if err != nil {
		panic("jwt middleware failed")
	}
	r.GET("/", homeLink)

	//users
	r.POST("/login", authMiddleware.LoginHandler)

	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	info.GET("/", a.HelloHandler)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	r.POST("/register", a.RegisterUser)

	// users
	admin.GET("/users", h.GetUsers)

	// products
	r.GET("/products/sorted", h.GetSortedProducts)
	r.GET("/products", h.GetProducts)
	r.GET("/product/:id", h.GetProduct)
	admin.POST("/product", h.AddProduct)
	admin.PUT("/product", h.UpdateProduct)
	admin.DELETE("/product", h.DeleteProduct)

	r.GET("/faq", h.GetFaq)
	admin.POST("/faq", h.AddFaq)
	admin.PUT("/faq", h.UpdateFaq)
	admin.DELETE("/faq", h.DeleteFaq)

	//categories
	r.GET("/categories", h.GetCategories)
	r.GET("/category", h.GetCategory)
	admin.POST("/category", h.AddCategory)
	admin.PUT("/category", h.UpdateCategory)
	admin.DELETE("/category", h.DeleteCategory)

	//images
	r.GET("/images", h.GetImages)
	r.GET("/image", h.GetImage)
	admin.POST("/image", h.AddImage)
	admin.DELETE("/image", h.DeleteImage)

	//orders
	order.GET("/:orderid", h.GetOrder)
	order.GET("/", h.GetOrders)
	order.POST("/", h.AddOrder)
	r.POST("/guest_order", h.AddGuestOrder)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	fmt.Println("Starting server...")
	r.Run(":8080")
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func CORS() cors.Config {
	return cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"X-Requested-With",
			"Access-Control-Allow-Origin",
			"Access-Control-Request-Headers",
			"Content-Type",
			"XMLHttpRequest",
			"Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}
}
