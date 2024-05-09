package controllers

import "github.com/gin-gonic/gin"

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (ac *AuthController) Login(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "Login Success",
	})

}
