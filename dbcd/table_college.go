package dbcd

import (
	"log"
)

// College 是表“College”的模型
type College struct {
	CollegeID   int
	CollegeName string
}

// CreateAndGetCollege 在数据表中创建College结构，并返回新的College结构的指针。
// 若College存在，则返回既有的结构。
func (engine *Engine) CreateAndGetCollege(name string) *College {
	query := `insert into "College" ("CollegeName") 
values (:1)`
	if _, err := engine.db.Exec(query, name); err != nil {
		log.Println(name, err)
	}
	return engine.GetCollegeByName(name)
}

// GetCollegeByName 返回name对应的College。
func (engine *Engine) GetCollegeByName(name string) *College {
	query := `select "CollegeID", "CollegeName" from "College" 
where "CollegeName"=:1`
	row := engine.db.QueryRow(query, name)
	
	var college College
	err := row.Scan(&college.CollegeID, &college.CollegeName)
	if err != nil {
		log.Println(name, err)
		return nil
	}
	return &college
}


// RemoveCollegeByName 从数据库中删除College结构。
func (engine *Engine) RemoveCollegeByName(name string) {
	query := `delete from "College" where "CollegeName"=:1`
	if _, err := engine.db.Exec(query, name); err != nil {
		log.Println(name, err)
	}
}


// TestCollege 测试College表。
func (engine *Engine) TestCollege() {
	log.Println("Test table College.")

	const testCollegeName = "只存在于传说中的学院"

	engine.CreateAndGetCollege(testCollegeName)
	engine.RemoveCollegeByName(testCollegeName)
}
