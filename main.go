package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	//logger
	r.Use(gin.Logger())

	r.GET("/hello", func(ctx *gin.Context) {
		ctx.IndentedJSON(200, map[string]any{
			"data": "test",
		})
	})

	r.Run(":8000")

	fmt.Println("Running on port : 8000")
}
