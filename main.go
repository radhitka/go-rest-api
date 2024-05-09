package main

import (
	"fmt"
	"go-rest-api/controllers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {

	r := gin.Default()

	//logger
	r.Use(gin.Logger())
	r.Use(testMiddleware)

	validate := validator.New()

	userContoller := controllers.NewAuthController(validate)

	api := r.Group("/api")

	api.POST("/login", userContoller.Login)

	r.Run(":8000")

	fmt.Println("Running on port : 8000")
}

func testMiddleware(c *gin.Context) {
	c.Next()
}
