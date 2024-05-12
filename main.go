package main

import (
	"fmt"
	"go-rest-api/config"
	"go-rest-api/controllers"
	"net/http"
	"os"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}

	db := config.SetUpDatabase()

	r := gin.Default()

	//logger
	r.Use(gin.Logger())

	validate := validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("form")
		if name == "" {
			name = fld.Name
		}
		return name
	})

	userContoller := controllers.NewAuthController(validate, db)
	bookController := controllers.NewBookController(validate, db)

	api := r.Group("/api")

	api.POST("/login", userContoller.Login)
	api.POST("/register", userContoller.Register)

	api.Use(authMiddleware)

	book := api.Group("/books")

	book.GET("/", bookController.GetBooks)
	book.POST("/add", bookController.AddBooks)

	r.Run("localhost:3000")

	fmt.Println("Running on port : 3000")
}

func authMiddleware(c *gin.Context) {

	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
		c.Abort()
		return
	}

	authHeader = authHeader[len("Bearer "):]

	token, err := jwt.Parse(authHeader, func(t *jwt.Token) (interface{}, error) {

		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unathorized",
		})
		c.Abort()
		return
	}

	c.Next()
}
