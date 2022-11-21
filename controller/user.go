package controller

import (
	"fmt"
	"go_demo/model"
	"go_demo/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	db := util.GetDB()

	//获取参数
	name := c.PostForm("name")
	mobile := c.PostForm("mobile")
	password := c.PostForm("password")
	userId := uuid.New().String()
	//数据验证
	if len(name) == 0 {
		util.WriteCustomResp(c, http.StatusUnprocessableEntity, "用户名不能为空")
		return
	}
	if len(mobile) != 11 {
		util.WriteCustomResp(c, http.StatusUnprocessableEntity, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		util.WriteCustomResp(c, http.StatusUnprocessableEntity, "密码不能少于6位")
		return
	}

	//判断手机号是否存在
	var user model.User
	db.Where("mobile = ?", mobile).First(&user)
	if user.ID != 0 {
		util.WriteCustomResp(c, http.StatusUnprocessableEntity, "用户已存在")
		return
	}

	//创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		util.WriteCustomResp(c, http.StatusInternalServerError, "密码加密错误")
		return
	}
	newUser := model.User{
		Name:     name,
		Mobile:   mobile,
		Password: string(hasedPassword),
		UserId:   userId,
	}
	db.Create(&newUser)

	//返回结果
	util.WriteSuccessResp(c, "注册成功")
}

type LoginParam struct {
	Mobile    string `json:"mobile"`    // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

func Login(c *gin.Context) {

	db := util.GetDB()

	//获取参数
	//此处使用Bind()函数，可以处理不同格式的前端数据
	var requestUser LoginParam
	err := c.ShouldBindJSON(&requestUser)
	if err != nil {
		util.WriteCustomResp(c, http.StatusBadRequest, "缺少必要参数")
		return
	}
	err = util.Verify(requestUser, util.LoginVerify)
	if err != nil {
		util.WriteCustomResp(c, http.StatusBadRequest, err.Error())
		return
	}
	if !(store.Verify(requestUser.CaptchaId, requestUser.Captcha, true)) {
		util.WriteCustomResp(c, http.StatusBadRequest, "验证码错误")
		return
	}
	//数据验证
	// if len(mobile) != 11 {
	// 	util.WriteCustomResp(c, http.StatusUnprocessableEntity, "手机号必须为11位")
	// 	return
	// }
	// if len(password) < 6 {
	// 	util.WriteCustomResp(c, http.StatusUnprocessableEntity, "密码不能少于6位")
	// 	return
	// }

	//判断手机号是否存在
	var user model.User
	db.Where("mobile = ?", requestUser.Mobile).First(&user)
	if user.ID == 0 {
		util.WriteCustomResp(c, http.StatusUnprocessableEntity, "用户不存在")
		return
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestUser.Password)); err != nil {
		util.WriteCustomResp(c, http.StatusUnprocessableEntity, "密码错误")
		return
	}
	tokenString, _ := util.GenToken(user.UserId, user.Mobile)
	fmt.Print("tokenString: ", tokenString)
	//返回结果
	c.Header("authorization", tokenString)
	util.WriteSuccessResp(c, "登录成功")
}
func LogOut(c *gin.Context) {
	c.Header("authorization", "")
	util.WriteSuccessResp(c, "登出成功")
}
