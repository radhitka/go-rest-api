package model

type Book struct {
	ID          int    `gorm:"column:id;primary_key;autoIncrement" json:"-"`
	Title       string `gorm:"column:title" json:"title"`
	TotalPages  int    `gorm:"column:total_pages" json:"total_pages"`
	Cover       string `gorm:"column:cover" json:"cover"`
	Author      string `gorm:"column:author" json:"author"`
	Publisher   string `gorm:"column:publisher" json:"publisher"`
	IsPublished bool   `gorm:"column:is_published" json:"is_published"`
}

func (b *Book) TableName() string {
	return "users"
}
