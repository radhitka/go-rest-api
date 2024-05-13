package utils

import (
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func NewFileName(filename string) string {
	fileEx := filepath.Ext(filename)
	return "upload_" + strconv.Itoa(int(time.Now().Unix())) + fileEx
}

func Env(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func GetJwtKey() string {
	return Env("JWT_KEY", "your-secret")

}
