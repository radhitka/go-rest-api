package config

import (
	"fmt"
	"go-rest-api/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetUpDatabase() *gorm.DB {

	DbHost := utils.Env("DB_HOST", "localhost")
	DbUser := utils.Env("DB_USER", "root")
	DbPassword := utils.Env("DB_PASSWORD", "")
	DbName := utils.Env("DB_NAME", "book_store")
	DbPort := utils.Env("DB_PORT", "3306")

	driver := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DbUser,
		DbPassword,
		DbHost,
		DbPort,
		DbName,
	)
	db, err := gorm.Open(mysql.Open(driver), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		panic(err)
	}

	return db

}
