package util

import (
	"fmt"
	"go_demo/config"
	"go_demo/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB(dbConfig *config.DatabaseConfig) *gorm.DB {
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DbName,
		dbConfig.Chartset)
	fmt.Printf("%+v\n", args)
	db, err := gorm.Open(dbConfig.Driver, args)
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}
	db.SingularTable(true)

	//迁移
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.UserInfo{})

	DB = db

	return db

}

func GetDB() *gorm.DB {
	return DB
}
