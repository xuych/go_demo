package util

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommonMessage struct {
	HttpCode int
	Msg      string
}

func (cm CommonMessage) WriteJsonResp(c *gin.Context) {
	c.JSON(cm.HttpCode, gin.H{
		"msg": cm.Msg,
	})
}

func WriteListResp(c *gin.Context, list interface{}, count int64, msg *string) {
	message := "获取列表成功"
	if msg != nil {
		message = *msg
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":   message,
		"data":  list,
		"total": count,
	})
}

func WriteErrResp(c *gin.Context, err error) {
	fmt.Printf("%+v\n", err)
	c.JSON(http.StatusInternalServerError, gin.H{
		"msg": err.Error(),
	})
}

func WriteCustomErrResp(c *gin.Context, httpCode int, err error) {
	fmt.Printf("%+v\n", err)
	c.JSON(httpCode, gin.H{
		"msg": err.Error(),
	})
}

var (
	Success     = CommonMessage{HttpCode: http.StatusOK, Msg: "success"}
	NoAuthHead  = CommonMessage{HttpCode: http.StatusUnauthorized, Msg: "auth header not found"}
	AuthFiled   = CommonMessage{HttpCode: http.StatusUnauthorized, Msg: "auth failed"}
	LoginFailed = CommonMessage{HttpCode: http.StatusUnauthorized, Msg: "username or password is incorrect"}
)

func MyFunc(arg *string) {
	fmt.Printf(*arg)
}
