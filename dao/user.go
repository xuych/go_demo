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
func (UserInfoDao) GetUserPageList(keyword *string, size int, page int, id *int) ([]model.UserInfo, int64) {
	var bonusApps []model.UserInfo
	query := util.DB.Model(&model.UserInfo{})
	fmt.Println("size: ", size)
	if keyword != nil {
		fmt.Println("keyword: ", keyword)
		query = query.Where("(name like ? OR id like ?)", keyword, keyword)
	}
	if id != nil {
		fmt.Println("keyword: ", keyword)
		query = query.Where("id = ?", id)
	}
	query.Order("id").
		Offset((page - 1) * size).
		Limit(size).Find(&bonusApps)

	var count int64
	query.Count(&count)

	return bonusApps, count
}
