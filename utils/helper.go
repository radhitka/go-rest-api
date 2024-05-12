package utils

import (
	"path/filepath"
	"strconv"
	"time"
)

func NewFileName(filename string) string {
	fileEx := filepath.Ext(filename)
	return "upload_" + strconv.Itoa(int(time.Now().Unix())) + fileEx
}
