package dbcd

import (
	"log"
)

// College 是表“College”的模型
type College struct {
	CollegeID   int
	CollegeName string
}

// CollegeExists 返回具有指定名称name的College是否存在。
func (engine *Engine) CollegeExists(name string) bool {
	query := `select count(*) from "College" where "CollegeName"=:1`
	result := engine.db.QueryRow(query, name)

	var count int
	if err := result.Scan(&count); err != nil {
		log.Println(query, name, err)
	}
	return count >= 1
}

// CreateCollege 创建并返回由参数指定的College结构的指针。
// 若College存在，则返回现有结构的指针。
func (engine *Engine) CreateCollege(name string) {
	query := `insert into "College" ("CollegeName") values (:1)`
	if _, err := engine.db.Exec(query, name); err != nil {
		log.Println(query, name, err)
	}
}

// GetCollegeByID 返回由ID指定的学院
func (engine *Engine) GetCollegeByID(id int) *College {
	query := `select "CollegeID", "CollegeName" from "College" where "CollegeID"=:1`
	result := engine.db.QueryRow(query, id)

	var college College
	if err := result.Scan(&college.CollegeID, &college.CollegeName); err != nil {
		log.Println(query, id, err)
		return nil
	}
	return &college
}

// GetCollegeByName 返回name对应的College。
func (engine *Engine) GetCollegeByName(name string) *College {
	query := `select "CollegeID", "CollegeName" from "College" where "CollegeName"=:1`
	row := engine.db.QueryRow(query, name)

	var college College
	err := row.Scan(&college.CollegeID, &college.CollegeName)
	if err != nil {
		return nil
	}
	return &college
}

// DeleteCollegeByName 从数据库中删除College结构。
func (engine *Engine) DeleteCollegeByName(name string) {
	college := engine.GetCollegeByName(name)
	if college == nil {
		return
	}
	id := college.CollegeID

	// 删除College的子记录
	queryDeleteChildRecords := `delete from "Specialty" where "CollegeID"=:1`
	if _, err := engine.db.Exec(queryDeleteChildRecords, id); err != nil {
		log.Println(queryDeleteChildRecords, id, err)
	}

	// 删除College
	query := `delete from "College" where "CollegeID"=:1`
	if _, err := engine.db.Exec(query, id); err != nil {
		log.Println(query, id, err)
	}
}

// TestTableCollege 测试College表。
func (engine *Engine) TestTableCollege() {
	log.Println("Testing table College.")

	const testCollegeName = "若此学院可见，数据表测试可能没有成功。"

	engine.DeleteCollegeByName(testCollegeName)
	engine.CreateCollege(testCollegeName)

	if !engine.CollegeExists(testCollegeName) {
		log.Panicln("Table College test failed! College with testCollegeName should exist.")
	}

	college := engine.GetCollegeByName(testCollegeName)
	if college == nil {
		log.Panicln("Table College test failed! college should NOT be nil.")
	}

	collegeID := college.CollegeID
	collegeName := college.CollegeName
	if collegeName != testCollegeName {
		log.Panicln("Table College test failed! CollegeName should be same with testCollegeName, not " + collegeName + " .")
	}
	if engine.GetCollegeByID(collegeID).CollegeName != testCollegeName {
		log.Panicln("Table College test failed! CollegeName should be same with testCollegeName, not " + collegeName + " .")
	}

	engine.DeleteCollegeByName(testCollegeName)

	if engine.CollegeExists(testCollegeName) {
		log.Panicln("Table College test failed! College with name testCollegeName should NOT exist.")
	}
}
