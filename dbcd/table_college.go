package dbcd

import (
	"log"
)

// College 是表“College”的模型
type College struct {
	CollegeID   int
	CollegeName string
}

// CreateAndGetCollege 创建并返回由参数指定的College结构的指针。
// 若College存在，则返回现有结构的指针。
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

// TestTableCollege 测试College表。
func (engine *Engine) TestTableCollege() {
	log.Println("Testing table College.")

	const testCollegeName = "若此学院可见，数据表测试可能没有成功。"

	engine.CreateAndGetCollege(testCollegeName)
	engine.RemoveCollegeByName(testCollegeName)
}
