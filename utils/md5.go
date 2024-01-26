package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

// 生成随机盐值
func GenerateSalt() string {
	randSource := rand.NewSource(time.Now().UnixNano())
	r := rand.New(randSource)

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	salt := make([]byte, 16)
	for i := range salt {
		salt[i] = charset[r.Intn(len(charset))]
	}
	return string(salt)
}

// 使用随机盐生成密码哈希值
func HashPassword(password string, salt string) string {
	// 将密码和盐值进行混合
	passwordWithSalt := password + salt

	// 使用 MD5 哈希算法生成摘要
	hash := md5.New()
	hash.Write([]byte(passwordWithSalt))
	hashInBytes := hash.Sum(nil)

	// 将摘要转换为十六进制字符串
	hashString := hex.EncodeToString(hashInBytes)

	return hashString
}

// 解密
func VerifyPassword(password string, salt string, dbPassword string) bool {
	passwordWithSalt := password + salt
	hash := md5.New()
	hash.Write([]byte(passwordWithSalt))
	hashInBytes := hash.Sum(nil)
	hashedPassword := hex.EncodeToString(hashInBytes)
	return hashedPassword == dbPassword
}
