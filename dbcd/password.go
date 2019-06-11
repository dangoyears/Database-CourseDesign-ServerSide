package dbcd

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

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
