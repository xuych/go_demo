package dao

import (
	"errors"
	"fmt"
	"go_demo/util"
)

type BaseDao struct {
}

// 新增
func (BaseDao) Add(model interface{}) (err error) {
	err = util.DB.Create(model).Error
	return
}

// 根据Id查询
func (BaseDao) Get(model interface{}, id int) error {
	// fmt.Println("UserInfoDao Get")
	fmt.Println("UserInfoDao Get")
	fmt.Println("UserInfoDao Get2")
	if id < 1 {
		return errors.New("请输入id")
	}
	err := util.DB.Where("id = ?", id).First(&model).Error
	return err
}

// 更新
func (BaseDao) UpdateModel(value interface{}) (err error) {
	err = util.DB.Save(value).Error
	return
}

// 删除
func (BaseDao) DeleteModel(value interface{}) (err error) {
	err = util.DB.Delete(value).Error
	return
}
