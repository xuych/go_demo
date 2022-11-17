package model

// 数据库表结构
type UserInfo struct {
	Id   int    `gorm:"primaryKey;autoIncrement" form:"id"`
	Name string `form:"name"`
	Age  int    `form:"age"`
}

// 自定义表名
func (UserInfo) TableName() string {
	return "userinfo"
}
