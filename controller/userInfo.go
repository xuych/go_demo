package controller

import (
	"go_demo/dao"
	"go_demo/model"
	"go_demo/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

// 新增用户
func (controller *UserController) Add(c *gin.Context) {
	//name := c.PostForm("name")
	var user model.UserInfo
	err := c.ShouldBindJSON(&user)
	if err != nil {
		util.WriteErrResp(c, err)
		return
	}
	dao.UserInfo.Add(&user)

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// 查询用户
func (controller *UserController) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	println("id:", id)
	var user model.UserInfo
	err := dao.UserInfo.Get(&user, id)
	if err != nil {
		util.WriteErrResp(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": &user,
	})
}

// 分页查询
func (controller *UserController) GetPageList(c *gin.Context) {
	param := util.GenPagination(c)
	// param := util.GenPagination(c)
	list, count := dao.UserInfo.GetUserPageList(&param)
	util.WriteListResp(c, list, count, nil)
}

// 更新name
func (controller *UserController) UpdateUserName(c *gin.Context) {
	name := c.PostForm("name")
	println("name:", name)
	id, _ := strconv.Atoi(c.PostForm("id"))
	println("id:", id)

	_ = dao.UserInfo.UpdateUserName(name, id)
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
	})
}
