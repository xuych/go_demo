package dao

import (
	"fmt"
	"go_demo/model"
	"go_demo/util"
)

type UserInfoDao struct {
	BaseDao
}

var UserInfo = UserInfoDao{}

// 修改用户名
func (UserInfoDao) UpdateUserName(name string, id int) error {
	err := util.DB.Model(&model.User{}).Where("id = ?", id).Update("name", name).Error
	return err
}

// 分页查询用户列表
func (UserInfoDao) GetUserPageList(Param *util.UserQueryParam) ([]model.UserInfo, int64) {
	fmt.Printf("Param: %+v", *Param)
	results := make([]model.UserInfo, 0)
	var count int64
	// util.DB.Model(&model.UserInfo{}).Count(&count)
	// util.DB.Model(&model.UserInfo{}).Limit(size).Offset(0).Find(&results)

	// util.DB.Model(&model.UserInfo{}).Limit(size).Offset((page - 1) * size).Find(&results)
	return results, count
}
