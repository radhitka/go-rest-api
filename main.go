package main

import (
	"fmt"
	"go-rest-api/config"
	"go-rest-api/controllers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	r.Use(testMiddleware)

	validate := validator.New()

	userContoller := controllers.NewAuthController(validate, db)

	api := r.Group("/api")

	api.POST("/login", userContoller.Login)
	api.POST("/register", userContoller.Register)

	r.Run(":8000")

	fmt.Println("Running on port : 8000")
}

func testMiddleware(c *gin.Context) {
	c.Next()
}
