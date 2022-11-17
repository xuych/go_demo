package model

import "github.com/jinzhu/gorm"

// 自定义表名
func (UserInfo) TableName() string {
	return "userinfo"
}

// 数据库表结构
type UserInfo struct {
	gorm.Model
	Name string `form:"name"`
	Age  int    `form:"age"`
}

type UserListQuery struct {
	Keyword *string `form:"keyword"`
	Page    int     `form:"page"`
	Size    int     `form:"size"`
	Id      *int    `form:"id"`
}
