package controllers

import (
	"go-rest-api/model"
	"go-rest-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type BookController struct {
	validate *validator.Validate
	DB       *gorm.DB
}

func NewBookController(validate *validator.Validate, db *gorm.DB) *BookController {
	return &BookController{
		validate: validate,
		DB:       db,
	}
}

func (bc *BookController) GetBooks(ctx *gin.Context) {
	var books []model.Book

	row, err := bc.DB.Raw("select id,title,total_pages,cover,author,publisher,is_published from books").Rows()

	res := utils.NewResponseData()

	if err != nil {

		res.WithMessage(err.Error()).InternalServerError()

		ctx.IndentedJSON(http.StatusInternalServerError, res)
		return
	}

	defer row.Close()

	for row.Next() {
		err = bc.DB.ScanRows(row, &books)

		if err != nil {

			res.WithMessage(err.Error()).InternalServerError()

			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
	}

	res.WithData(books).SuccessOk()

	ctx.IndentedJSON(http.StatusOK, res)

}
