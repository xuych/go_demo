package common

import (
	"fmt"
	"go_demo/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "127.0.0.1"
	port := "3306"
	database := "go_demo"
	username := "root"
	password := "1234"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}
	db.SingularTable(true)

	//迁移
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.App{})

	DB = db

	return db

}

func GetDB() *gorm.DB {
	return DB
}
