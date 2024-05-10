package controllers

import (
	"go-rest-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/go-playground/validator/v10"
)

type AuthController struct {
	validate *validator.Validate
	DB       *gorm.DB
}

func NewAuthController(validator *validator.Validate, db *gorm.DB) *AuthController {
	return &AuthController{
		validate: validator,
		DB:       db,
	}
}

type LoginRequest struct {
	Username string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
}
type Register struct {
	Name     string `form:"name" validate:"required"`
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

func (ac *AuthController) Register(ctx *gin.Context) {

	var registerRequest Register

	ctx.Bind(&registerRequest)

	err := ac.validate.Struct(registerRequest)

	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var user model.User

	err = ac.DB.Take(&user, "username = ?", registerRequest.Username).Error

	if err != nil && err != gorm.ErrRecordNotFound {

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if user.Username != "" {

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Username sudah digunakan",
		})
		return
	}

	newPassword := []byte(registerRequest.Password)

	hashedPassword, err := bcrypt.GenerateFromPassword(newPassword, bcrypt.DefaultCost)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	user = model.User{
		Name:     registerRequest.Name,
		Username: registerRequest.Username,
		Password: string(hashedPassword),
	}

	err = ac.DB.Create(&user).Error

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Register Success",
	})

}
