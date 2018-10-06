package db

import (
	"github.com/jinzhu/gorm"
	"os"
)

func GetMysql() *gorm.DB {
	db, _ := gorm.Open("mysql", os.Getenv("MYSQL_STRING_CONNECT"))
	return db
}
