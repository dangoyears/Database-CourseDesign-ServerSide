package dbcd

import "log"

// Administrator 表“Administrator”的抽象
type Administrator struct {
	AdminLoginName string
	AdminPassHash  string
}

// CreateAdministrator 使用用户名和密码来创建管理员账户。
func (engine *Engine) CreateAdministrator(user, pass string) {
	passHash := GeneratePasswordHash(pass)

	query := `insert into "Administrator" ("AdminLoginName", "AdminPassHash")
values (:1, :2)`
	_, err := engine.db.Exec(query, user, passHash)
	if err != nil {
		Trace(err, query, user)
	}
}

// CreateDefaultAdministrator 使用默认用户名和密码创建管理员账户。
func (engine *Engine) CreateDefaultAdministrator() {
	const (
		username = "dangoyears"
		password = "dangoyears"
	)
	engine.CreateAdministrator(username, password)
}

// GetAdministratorsCount 返回系统中所有管理员的数量。
func (engine *Engine) GetAdministratorsCount() int {
	query := `select count(*) from "Administrator"`
	result := engine.db.QueryRow(query)

	var count int
	result.Scan(&count)
	return count
}

// ExistAdministratorWithName 返回具有name登陆名的管理员账户是否存在。
func (engine *Engine) ExistAdministratorWithName(name string) bool {
	query := `select count(*) from "Administrator where "AdminLoginName"=:1`
	result := engine.db.QueryRow(query, name)

	var count int
	if err := result.Scan(&count); err != nil {
		Trace(err, query, name)
	}
	return count >= 1
}

// GetAdministratorByLoginName 返回具有LoginName的Administrator。
func (engine *Engine) GetAdministratorByLoginName(name string) *Administrator {
	var admin Administrator
	query := `select "AdminLoginName", "AdminPassHash" from "Administrator" where "AdminLoginName"=:1`
	row := engine.db.QueryRow(query, name)
	if err := row.Scan(&admin.AdminLoginName, &admin.AdminPassHash); err != nil {
		return nil
	}
	return &admin
}

// DeleteAdministratorByLoginName 删除指定登陆名称的管理员。
func (engine *Engine) DeleteAdministratorByLoginName(name string) {
	query := `delete from "Administrator" where "AdminLoginName"=:1`
	_, err := engine.db.Exec(query, name)
	if err != nil {
		Trace(err, query, name)
	}
}

// TestTableAdministrator 测试表Administrator。
func (engine *Engine) TestTableAdministrator() {
	log.Println("Testing table Administrator.")

	// 准备测试环境。
	const (
		testName = "（测试管理员账户名）"
		testPass = "（测试密码）"
	)
	engine.DeleteAdministratorByLoginName(testName)

	// 测试。
	countBeforeCreate := engine.GetAdministratorsCount()
	engine.CreateAdministrator(testName, testPass)
	countAfterCreate := engine.GetAdministratorsCount()
	if countAfterCreate != countBeforeCreate+1 {
		log.Panicln("Table Administrator test failed! countAfterCreate should be greater than countBeforeCreate by one.")
	}
	if engine.GetAdministratorByLoginName(testName) == nil {
		log.Panicln("Table Administrator test failed! Should be able to getAdministratorbyLoginName.")
	}
	engine.DeleteAdministratorByLoginName(testName)
	if engine.GetAdministratorByLoginName(testName) != nil {
		log.Panicln("Table Administrator test failed! Should NOT be able to getAdministratorbyLoginName.")
	}
}
