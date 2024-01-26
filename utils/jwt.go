package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	JwtKey             = "sys-admin"                                         //密钥
	TokenExpire        = time.Now().Add(time.Second * 3600 * 24 * 7).Unix()  // 有效期 7
	RefreshTokenExpire = time.Now().Add(time.Second * 3600 * 24 * 14).Unix() // 刷新有效期 14
)

type UserClaim struct {
	Id      uint
	Name    string
	IsAdmin bool // 管理
	jwt.StandardClaims
}


// 生成token
func GenerateToken(id uint, name string, expireAt int64) (string, error) {
	user := UserClaim{
		Id:   id,
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	userToken, err := token.SignedString([]byte(JwtKey))

	if err != nil {
		return "", err
	}
	return userToken, nil
}

