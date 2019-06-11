package dbcd

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// GenerateToken 返回一个随机生成的32位长度的token。
func GenerateToken() string {
	rand.Seed(rand.Int63() + time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	tokenLength := 32
	var tokenBuilder strings.Builder
	for i := 0; i < tokenLength; i++ {
		tokenBuilder.WriteRune(chars[rand.Intn(len(chars))])
	}
	token := tokenBuilder.String()
	return token
}

// GeneratePasswordHash 返回哈希后的密码。
func GeneratePasswordHash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln(err)
	}
	return string(hash)
}

// MatchPasswordAndHash 比较密码与密码哈希是否相符。
func MatchPasswordAndHash(password, hash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}
	return true
}
