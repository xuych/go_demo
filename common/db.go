package common

import (
	"fmt"
	"go_demo/config"
	"go_demo/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dbConfig := config.GetConfig().Database
	driverName := dbConfig.Driver
	host := dbConfig.Host
	port := dbConfig.Port
	database := dbConfig.DbName
	username := dbConfig.User
	password := dbConfig.Password
	charset := dbConfig.Chartset
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	fmt.Printf("%+v\n", args)
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
