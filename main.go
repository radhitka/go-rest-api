package main

import (
	"fmt"
	"go-rest-api/config"
	"go-rest-api/controllers"
	"go-rest-api/utils"
	"net/http"
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
	book.GET("/detail/:id", bookController.DetailBooks)
	book.POST("/add", bookController.AddBooks)
	book.POST("/update/:id", bookController.UpdateBooks)
	book.DELETE("/delete/:id", bookController.DeleteBook)

	port := utils.Env("PORT", "3000")

	r.Run(fmt.Sprintf("localhost:%s", port))

	fmt.Printf("Running on port : %s", port)
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

		return []byte(utils.GetJwtKey()), nil
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
