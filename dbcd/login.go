package dbcd

import "strconv"

// LoginAdmin 若传入的用户名和密码正确则返回token，
// 否则返回空字符串。
func (engine *Engine) LoginAdmin(name, pass string) string {
	admin := engine.GetAdministratorByLoginName(name)
	if admin != nil && MatchPasswordAndHash(pass, admin.AdminPassHash) {
		token := GenerateToken()
		engine.keeper.addRoleToken(token, "admin")
		return token
	}
	return ""
}

// LoginStudent 若传入的用户名和密码正确则返回token，
// 否则返回空字符串。
func (engine *Engine) LoginStudent(name, pass string) string {
	studentNumber, studentNumberErr := strconv.Atoi(name)

	if studentNumberErr != nil {
		return ""
	}

	student := engine.GetStudentInfoByStudentNumber(studentNumber)
	if student == nil {
		return ""
	}

	if MatchPasswordAndHash(pass, student.PasswordHash) {
		token := GenerateToken()
		engine.keeper.addRoleToken(token, "student")
		return token
	}
	return ""
}

// LoginTeacher 若传入的用户名和密码正确则返回token，
// 否则返回空字符串。
func (engine *Engine) LoginTeacher(name, pass string) string {
	teacherNumber, teacherNumberErr := strconv.Atoi(name)

	if teacherNumberErr != nil {
		return ""
	}

	teacher := engine.GetTeacherInfoByTeacherNumber(teacherNumber)
	if teacher == nil {
		return ""
	}

	if MatchPasswordAndHash(pass, teacher.PasswordHash) {
		token := GenerateToken()
		engine.keeper.addRoleToken(token, "teacher")
		return token
	}
	return ""
}
