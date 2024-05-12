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

	for i := 0; i < len(books); i++ {
		books[i].Cover = "assets/images/" + books[i].Cover
	}

	res.WithData(books).SuccessOk()

	ctx.IndentedJSON(http.StatusOK, res)

}

func (bc *BookController) AddBooks(ctx *gin.Context) {
	bookRequest := model.Book{}
	err := ctx.Bind(&bookRequest)

	res := utils.NewResponseData()
	if err != nil {
		res.WithMessage(err.Error()).BadRequest()

		ctx.JSON(res.StatusCode, res)
		return
	}

	err = bc.validate.Struct(bookRequest)

	if err != nil {
		res.WithMessage(err.Error()).BadRequest()

		ctx.JSON(res.StatusCode, res)
		return
	}

	cover, err := ctx.FormFile("cover")

	if err != nil {
		res.WithMessage(err.Error()).BadRequest()

		ctx.JSON(res.StatusCode, res)
		return
	}

	newFileName := utils.NewFileName(cover.Filename)

	path := "assets/images/" + newFileName
	err = ctx.SaveUploadedFile(cover, path)

	if err != nil {
		res.WithMessage(err.Error()).InternalServerError()

		ctx.JSON(res.StatusCode, res)
		return
	}

	bookRequest.Cover = newFileName

	err = bc.DB.Create(&bookRequest).Error

	if err != nil {
		res.WithMessage(err.Error()).InternalServerError()

		ctx.JSON(res.StatusCode, res)
		return
	}

	res.SuccessCreated()
	ctx.JSON(res.StatusCode, res)
}

func (bc *BookController) DetailBooks(ctx *gin.Context) {

	id := ctx.Param("id")

	var book model.Book

	res := utils.NewResponseData()

	err := bc.DB.First(&book, "id = ?", id).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		res.WithMessage(err.Error()).InternalServerError()

		ctx.JSON(res.StatusCode, res)
		return
	}

	if err == gorm.ErrRecordNotFound {
		res.WithMessage("Buku tidak ditemukan!").NotFound()

		ctx.JSON(res.StatusCode, res)
		return
	}

	book.Cover = "asset/images/" + book.Cover

	res.WithData(book).SuccessOk()

	ctx.JSON(res.StatusCode, res)
}
