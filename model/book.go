package model

import "mime/multipart"

type Book struct {
	ID          int                   `gorm:"column:id;primary_key;autoIncrement" json:"-"`
	Title       string                `gorm:"column:title" json:"title" form:"title" validate:"required"`
	TotalPages  int                   `gorm:"column:total_pages" json:"total_pages" form:"total_pages" validate:"required"`
	Cover       string                `gorm:"column:cover" json:"cover"`
	CoverUpload *multipart.FileHeader `gorm:"-" json:"-" form:"cover" validate:"required"`
	Author      string                `gorm:"column:author" json:"author" form:"author" validate:"required"`
	Publisher   string                `gorm:"column:publisher" json:"publisher" form:"publisher" validate:"required"`
	IsPublished bool                  `gorm:"column:is_published" json:"is_published" form:"is_published" validate:"required"`
}

func (b *Book) TableName() string {
	return "books"
}
