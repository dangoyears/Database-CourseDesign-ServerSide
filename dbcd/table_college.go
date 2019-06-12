package dbcd

// College 是表“College”的模型
type College struct {
	ID   int
	Name string
}

// GetCollegeByName 返回name对应的College
func (engine *Engine) GetCollegeByName(name string) *College {
	query := `select "CollegeID", "CollegeName" from "College" 
where "CollegeName"=:1`
	row := engine.db.QueryRow(query, name)
	var college College
	if err := row.Scan(&college.ID, college.Name); err == nil {
		return &college
	}
	return nil
}
