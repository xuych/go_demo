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
	UserId     string `json:"userId"`
	Mobile     string `json:"mobile"`
	BufferTime int64
	jwt.StandardClaims
}
type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("令牌已过期，请重新登录")
	TokenNotValidYet = errors.New("令牌尚未激活")
	TokenMalformed   = errors.New("令牌无效")
	TokenInvalid     = errors.New("令牌无效:")
)

const TokenExpireDuration = time.Second * 60 * 60 * 24 // token过期时间 m
const TokenBufferDuration = time.Second * 60 * 60 * 24 // token自动更新的缓冲时间 m

// const TokenExpiresTime = time.Second * 30

var Secret = []byte("secret") // 密码自行设定

func GenToken(uid string, mobile string) (string, error) {
	c := MyClaims{
		uid,
		mobile,
		time.Now().Add(TokenBufferDuration).Unix(),
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
func ParseToken(token string, c *gin.Context) (*MyClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// 验证是否过期且在有buffer内，是的话更新token到header
				if claims, ok := tokenClaims.Claims.(*MyClaims); ok {
					if claims.BufferTime > time.Now().Unix() {
						newTokenString, _ := GenToken(claims.Mobile, claims.UserId)
						c.Header("authorization", newTokenString)
						return nil, nil
					}
					return nil, TokenExpired
				}

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
		// http.StatusUnauthorized
		WriteCustomErrResp(c, http.StatusUnauthorized, errors.New("令牌不能为空"))
		c.Abort()
		return
	}
	// 按空格分割
	parts := strings.Split(authHeader, ".")
	if len(parts) != 3 {
		WriteCustomErrResp(c, http.StatusUnauthorized, errors.New("令牌格式有误"))
		c.Abort()
		return
	}

	_, err := ParseToken(authHeader, c)
	if err != nil {
		WriteErrResp(c, err)
		c.Abort()
		return
	}
	c.Next() // 后续的处理函数可以用过c.Get("mobile")来获取当前请求的用户信息
}
