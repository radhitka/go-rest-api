package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/go-playground/validator/v10"
)

type AuthController struct {
	validate *validator.Validate
}

func NewAuthController(validator *validator.Validate) *AuthController {
	return &AuthController{
		validate: validator,
	}
}

type LoginRequest struct {
	Username string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
}

func (ac *AuthController) Login(ctx *gin.Context) {

	var loginRequest LoginRequest

	ctx.Bind(&loginRequest)

	err := ac.validate.Struct(loginRequest)

	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	username := loginRequest.Username

	ctx.JSON(200, gin.H{
		"data":    username,
		"message": "Login Success",
	})

}
