package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"varchar(20);not null"`
	Mobile   string `gorm:"varchar(20);not null;unique"`
	Password string `gorm:"size:255;not null"`
	UserId   string // 用户唯一ID
	RoleId   int    // 权限类型
}
