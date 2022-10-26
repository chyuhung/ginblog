package middleware

import (
	"ginblog/utils"
	"github.com/golang-jwt/jwt"
	"time"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// SetToken 生成token
func SetToken(username string, password string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour)
	SetClaims := MyClaims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ginblog",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodES256, SetClaims)
}

// 验证token

// jwt中间件
