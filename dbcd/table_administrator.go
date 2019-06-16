package dbcd

import "log"

// Administrator 表“Administrator”的抽象
type Administrator struct {
	LoginName string
	PassHash  string
}

// GetAdministratorsCount 返回系统中所有管理员的数量。
func (engine *Engine) GetAdministratorsCount() int {
	query := `select count(*) from "Administrator"`
	result := engine.db.QueryRow(query)

	var count int
	result.Scan(&count)
	return count
}

// AdministratorExists 返回具有name登陆名的管理员账户是否存在
func (engine *Engine) AdministratorExists(name string) bool {
	query := `select count(*) from "Administrator where "AdminLoginName"=:1`
	result := engine.db.QueryRow(query, name)

	var count int
	if err := result.Scan(&count); err != nil {
		log.Println(name, err)
	}
	return count >= 1
}

// CreateDefaultAdministrator 使用默认用户名和密码创建管理员账户。
func (engine *Engine) CreateDefaultAdministrator() {
	const (
		username = "dangoyears"
		password = "dangoyears"
	)
	engine.CreateAdministrator(username, password)
}

// CreateAdministrator 使用用户名和密码来创建管理员账户。
func (engine *Engine) CreateAdministrator(user, pass string) {
	passHash := GeneratePasswordHash(pass)

	query := `insert into "Administrator" ("AdminLoginName", "AdminPassHash")
values (:1, :2)`
	_, err := engine.db.Exec(query, user, passHash)
	if err != nil {
		log.Println(user, err)
	}
}

// GetAdministratorByLoginName 返回具有LoginName的Administrator。
func (engine *Engine) GetAdministratorByLoginName(name string) *Administrator {
	var admin Administrator
	query := `select "AdminLoginName", "AdminPassHash" from "Administrator" where "AdminLoginName"=:1`
	row := engine.db.QueryRow(query, name)
	if err := row.Scan(&admin.LoginName, &admin.PassHash); err != nil {
		return nil
	}
	return &admin
}

// DeleteAdministratorByLoginName 删除指定登陆名称的管理员。
func (engine *Engine) DeleteAdministratorByLoginName(name string) {
	query := `delete from "Administrator" where "AdminLoginName"=:1`
	_, err := engine.db.Exec(query, name)
	if err != nil {
		log.Println(name, err)
	}
}

// TestTableAdministrator 测试表Administrator。
func (engine *Engine) TestTableAdministrator() {
	log.Println("Testing table Administrator.")

	const (
		testName = "如果此用户名可见，表测试可能没有成功。"
		testPass = "This is a password for test."
	)

	engine.DeleteAdministratorByLoginName(testName)

	countBeforeCreate := engine.GetAdministratorsCount()
	engine.CreateAdministrator(testName, testPass)
	countAfterCreate := engine.GetAdministratorsCount()
	if countAfterCreate != countBeforeCreate+1 {
		log.Panicln("Table Administrator test failed! countAfterCreate should be greater than countBeforeCreate by one.")
	}
	if engine.GetAdministratorByLoginName(testName) == nil {
		log.Panicln("Table Administartor test failed! Should be able to getAdministratorbyLoginName.")
	}

	engine.DeleteAdministratorByLoginName(testName)
	if engine.GetAdministratorByLoginName(testName) != nil {
		log.Panicln("Table Administartor test failed! Should NOT be able to getAdministratorbyLoginName.")
	}
}
