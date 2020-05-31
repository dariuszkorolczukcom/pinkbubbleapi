package auth

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	db "github.com/dariuszkorolczukcom/pinkbubbleapi/database"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

var identityKey = "ID"

func RegisterUser(c *gin.Context) {
	var err error
	var credentials User
	fmt.Println(c)
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), 8)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	credentials.Password = string(hashedPassword)
	db.Conn.Create(&credentials)
	c.JSON(200, credentials)
}

func AuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte("secret key"),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     identityKey,
		PayloadFunc:     payloadFunc,
		IdentityHandler: identityHandlerFunc,
		Authenticator:   authenticatorFunc,
		Authorizator:    authorizatorFunc,
		Unauthorized:    unauthorizedFunc,
	})

	return authMiddleware, err
}

func AuthMiddlewareClient() (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte("secret key"),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     identityKey,
		PayloadFunc:     payloadFunc,
		IdentityHandler: identityHandlerFunc,
		Authenticator:   authenticatorFunc,
		Authorizator:    authorizatorFuncClient,
		Unauthorized:    unauthorizedFunc,
	})

	return authMiddleware, err
}

func IsAdmin(c *gin.Context) bool {
	claims := jwt.ExtractClaims(c)
	fmt.Println(claims)
	user, _ := c.Get(identityKey)
	fmt.Println(user)
	if user.(User).Role == 1 {
		return true
	}
	return false
}

func HelloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	fmt.Println(user)
	c.JSON(200, gin.H{
		"userID":    claims[identityKey],
		"email":     user.(User).Email,
		"firstName": user.(User).FirstName,
		"lastName":  user.(User).LastName,
		"role":      user.(User).Role,
	})
}

func payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(User); ok {
		return jwt.MapClaims{
			identityKey: v.ID,
		}
	}
	return jwt.MapClaims{}
}

func identityHandlerFunc(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	var user User
	db.Conn.Where("id = ?", claims[identityKey]).First(&user)
	user.Password = ""
	return user
}

func authenticatorFunc(c *gin.Context) (interface{}, error) {
	var err error
	var credentials User
	var stored User
	if err := c.ShouldBindJSON(&credentials); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	db.Conn.Where(User{Email: credentials.Email}).First(&stored)
	if err = bcrypt.CompareHashAndPassword([]byte(stored.Password), []byte(credentials.Password)); err != nil {
		return nil, jwt.ErrFailedAuthentication
	}
	stored.Password = ""
	return stored, nil
}

func authorizatorFunc(data interface{}, c *gin.Context) bool {
	if v, ok := data.(User); ok && v.Role == 1 {
		return true
	}
	return false
}

func authorizatorFuncClient(data interface{}, c *gin.Context) bool {
	return true
}

func unauthorizedFunc(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

type User struct {
	gorm.Model
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      int    `json:"status"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func GetUserId(c *gin.Context) float64 {
	id := jwt.ExtractClaims(c)
	return id["ID"].(float64)
}
