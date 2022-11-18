package util

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

type MyClaims struct {
	UserId      string `json:"userId"`
	Mobile      string `json:"mobile"`
	BufferTime  int64
	ExpiresTime int64
	jwt.StandardClaims
}

var (
	TokenExpired     = errors.New("令牌已过期，请重新登录")
	TokenNotValidYet = errors.New("令牌尚未激活")
	TokenMalformed   = errors.New("令牌无效")
	TokenInvalid     = errors.New("令牌无效:")
)

const TokenExpireDuration = time.Second * 30 // 过期时间
const TokenBufferTime = 60 * 60 * 24 * 1
const TokenExpiresTime = 60 * 60 * 24 * 7

var Secret = []byte("secret") // 密码自行设定
// var Secret = config.GlobalConfig.JwtConfig.JwtSecretKey // 密码自行设定

func GenToken(mobile string, uid string) (string, error) {
	c := MyClaims{
		uid,
		mobile,
		TokenBufferTime,
		TokenExpiresTime,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "go_demo",                                  // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(Secret)
}

// ParseToken 解析JWT
func ParseToken(token string) (*MyClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}

	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*MyClaims); ok && tokenClaims.Valid {
			return claims, nil
		}

	}
	return nil, err
}

// JWTAuthMiddleware 基于JWT的认证中间件--验证用户是否登录
func JwtAuthMiddleware(c *gin.Context) {
	authHeader := c.Request.Header.Get("authorization")
	fmt.Print("authHeader: ", authHeader)
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
		WriteErrResp(c, err)
		// CommonMessage.WriteJsonResp(AuthFiled, c)
		c.Abort()
		return
	}
	c.Next() // 后续的处理函数可以用过c.Get("mobile")来获取当前请求的用户信息
}
