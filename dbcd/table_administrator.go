package dbcd

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
	engine.db.Exec(query, user, passHash)
}

// GetAdministratorByLoginName 返回具有LoginName的Administrator。
func (engine *Engine) GetAdministratorByLoginName(name string) *Administrator {
	var admin Administrator
	query := `select "AdminLoginName", "AdminPassHash" from "Administrator" where "AdminLoginName"=:1`
	row := engine.db.QueryRow(query, name)
	if err := row.Scan(&admin.LoginName, &admin.PassHash); err == nil {
		return &admin
	}
	return nil
}
