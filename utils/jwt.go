package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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

// 验证token
func VerifyToken(token string) (*UserClaim, int) {

	setToken, err := jwt.ParseWithClaims(token, &UserClaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(JwtKey), nil
	})

	if err != nil {
		fmt.Println("Token解析失败:", err)
		return nil, 401
	}

	if claims, _ := setToken.Claims.(*UserClaim); setToken.Valid {
		return claims, 200
	} else {
		return nil, 401
	}
}

// JWT 中间件
func JwtToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenHeader := ctx.Request.Header.Get("Authorization")
		Result := &Result{}

		if tokenHeader == "" {
			Result.Message = "token 不存在"
			ctx.JSON(200, Result.Fail(
				401, Result.Message,
			))
			ctx.Abort()
			return
		}

		verifyToken := strings.SplitN(tokenHeader, " ", 2)

		if len(verifyToken) != 2 && verifyToken[0] != "Bearer" {
			Result.Message = "token 格式错误"
			ctx.JSON(200, Result.Fail(
				401, Result.Message,
			))
			ctx.Abort()
			return
		}

		key, code := VerifyToken(verifyToken[1])

		if code == 401 {
			Result.Message = "token 验证失败"
			ctx.JSON(200, Result.Fail(
				401, Result.Message,
			))
			ctx.Abort()
			return

		}

		if time.Now().Unix() > key.ExpiresAt {
			Result.Message = "token 过期"
			ctx.JSON(200, Result.Fail(
				401, Result.Message,
			))
			ctx.Abort()
			return
		}

		ctx.Set("name", key.Name)
		ctx.Next()
	}
}
