package dbcd

import (
	"math/rand"
	"strings"
)

// GateKeeper 保存用户登陆时的凭证
type GateKeeper struct {
	Token2HumanID map[string]string

	// LoginType 表示用户登陆的类型，为{"admin", "student", "teacher"}之一。
	Token2LoginType map[string]string
}

func NewGateKeeper() *GateKeeper {
	var keeper GateKeeper
	keeper.Token2HumanID = make(map[string]string)
	keeper.Token2LoginType = make(map[string]string)
	return &keeper
}

// GenerateToken 返回一个随机生成的32位长度的token
func (keeper *GateKeeper) GenerateToken() string {
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	tokenLength := 32
	var tokenBuilder strings.Builder
	for i := 0; i < tokenLength; i++ {
		tokenBuilder.WriteRune(chars[rand.Intn(len(chars))])
	}
	token := tokenBuilder.String()
	return token
}

// LoginAdmin 若传入的用户名和密码正确则返回token
// 否则返回空字符串
func (keeper *GateKeeper) LoginAdmin(name, pass string) string {
	if name == "dangoyears" && pass == "dangoyears" { // 硬编码用户名和密码
		token := keeper.GenerateToken()
		keeper.addTokenForLoginType(token, "admin")
		return token
	}
	return ""
}

// LoginStudent 若传入的用户名和密码正确则返回token
// 否则返回空字符串
// @未完成
func (keeper *GateKeeper) LoginStudent(name, pass string) string {
	return ""
}

// LoginTeacher 若传入的用户名和密码正确则返回token
// 否则返回空字符串
// @未完成
func (keeper *GateKeeper) LoginTeacher(name, pass string) string {
	return ""
}

// GenerateAndValidTokenForHumanLogin 返回一个有效的token给调用者
// 返回的token可用于认证自然人登陆
// @未完成
func (keeper *GateKeeper) GenerateAndValidTokenForHumanLogin() string {
	token := keeper.GenerateToken()
	return token
}

// RemoveTokenForHumanLogin 将传入的token失效
// @未完成
func (keeper *GateKeeper) RemoveTokenForHumanLogin(token string) {

}

func (keeper *GateKeeper) addTokenForHumanID(token, humanID string) {
	keeper.Token2HumanID[token] = humanID
}

func (keeper *GateKeeper) addTokenForLoginType(token, loginType string) {
	keeper.Token2LoginType[token] = loginType
}

func (keeper *GateKeeper) deleteToken(token string) {
	delete(keeper.Token2HumanID, token)
	delete(keeper.Token2LoginType, token)
}
