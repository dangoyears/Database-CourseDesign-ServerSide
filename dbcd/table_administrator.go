package dbcd

// Administrator 表“Administrator”的抽象
type Administrator struct {
	LoginName string
	PassHash  string
}

// GetAdministratorByLoginName 返回具有LoginName的Administrator
func (engine *Engine) GetAdministratorByLoginName(name string) *Administrator {
	var admin Administrator
	query := `select "AdminLoginName", "AdminPassHash" from "Administrator" where "AdminLoginName"=:1`
	row := engine.db.QueryRow(query, name)
	if err := row.Scan(&admin.LoginName, &admin.PassHash); err == nil {
		return &admin
	}
	return nil
}
