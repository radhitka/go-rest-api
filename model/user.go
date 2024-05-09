package model

type User struct {
	ID       int    `gorm:"column:id;primary_key;autoIncrement" json:"-"`
	Username string `gorm:"column:username" json:"username"`
	Name     string `gorm:"column:name" json:"name"`
	Password string `gorm:"column:password" json:"-"`
}

func (u *User) TableName() string {
	return "users"
}
