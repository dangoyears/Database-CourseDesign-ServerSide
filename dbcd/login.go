package dbcd

// LoginAdmin 若传入的用户名和密码正确则返回token，
// 否则返回空字符串。
func (engine *Engine) LoginAdmin(name, pass string) string {
	admin := engine.GetAdministratorByLoginName(name)
	if admin.LoginName == name && MatchPasswordAndHash(pass, admin.PassHash) { // 硬编码用户名和密码
		token := GenerateToken()
		engine.keeper.addRoleToken(token, "admin")
		return token
	}
	return ""
}

// LoginStudent 若传入的用户名和密码正确则返回token，
// 否则返回空字符串。
// @未完成
func (engine *Engine) LoginStudent(name, pass string) string {
	return ""
}

// LoginTeacher 若传入的用户名和密码正确则返回token，
// 否则返回空字符串。
// @未完成
func (engine *Engine) LoginTeacher(name, pass string) string {
	return ""
}