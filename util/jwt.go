package util

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

type RegisteredClaims struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

const TokenExpireDuration = time.Hour * 24 // 过期时间

var Secret = []byte("go_demo") // 密码自行设定

func GenToken(telephone string, uid string) (string, error) {
	// 创建一个我们自己的声明
	c := RegisteredClaims{
		telephone, // 自定义字段
		uid,       // 自定义字段
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)), // 过期时间
			Issuer:    "go_demo",                                               // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(Secret)
}

// ParseToken 解析JWT
func ParseToken(token string) (*RegisteredClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*RegisteredClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// JWTAuthMiddleware 基于JWT的认证中间件--验证用户是否登录
func JwtAuthMiddleware(c *gin.Context) {
	authHeader := c.Request.Header.Get("authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 2003,
			"msg":  "token 不能为空",
		})
		c.Abort()
		return
	}
	// 按空格分割
	parts := strings.Split(authHeader, ".")
	if len(parts) != 3 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 2004,
			"msg":  "token 格式有误",
		})
		c.Abort()
		return
	}
	_, err := ParseToken(authHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 2005,
			"msg":  "无效的Token",
		})
		c.Abort()
		return
	}
	c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
}
