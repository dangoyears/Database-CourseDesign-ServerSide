package dbcd

import (
	"log"
)

// College 是表“College”的模型
type College struct {
	CollegeID   int
	CollegeName string
}

// CreateCollege 创建并返回由参数指定的College结构的指针。
// 若College存在，则返回现有结构的指针。
func (engine *Engine) CreateCollege(name string) {
	query := `insert into "College" ("CollegeName") values (:1)`
	if _, err := engine.db.Exec(query, name); err != nil {
		Trace(err, query, name)
	}
}

// ExistCollege 返回具有指定名称name的College是否存在。
func (engine *Engine) ExistCollege(name string) bool {
	query := `select count(*) from "College" where "CollegeName"=:1`
	result := engine.db.QueryRow(query, name)

	var count int
	if err := result.Scan(&count); err != nil {
		Trace(err, query, name)
	}
	return count >= 1
}

// GetCollegeByID 返回由ID指定的学院。
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

// DeleteCollegeByName 从数据库中删除College项。
func (engine *Engine) DeleteCollegeByName(name string) {
	college := engine.GetCollegeByName(name)
	if college == nil {
		return
	}
	id := college.CollegeID

	// 删除College名下的教师。
	queryDeleteRelatedTeachers := `delete from "Teacher" where "CollegeID"=:1`
	if _, err := engine.db.Exec(queryDeleteRelatedTeachers, id); err != nil {
		Trace(err, queryDeleteRelatedTeachers, id)
	}

	// 删除College名下的专业和班级。
	queryDeleteRelatedClasses := `delete from "Class" where "SpecialtyID" in (select "SpecialtyID" from "Specialty" where "CollegeID"=:1)`
	queryDeleteRelatedSpecialties := `delete from "Specialty" where "CollegeID"=:1`
	if _, err := engine.db.Exec(queryDeleteRelatedClasses, id); err != nil {
		Trace(err, queryDeleteRelatedClasses, id)
	}
	if _, err := engine.db.Exec(queryDeleteRelatedSpecialties, id); err != nil {
		Trace(err, queryDeleteRelatedSpecialties, id)
	}

	// 删除College。
	query := `delete from "College" where "CollegeID"=:1`
	if _, err := engine.db.Exec(query, id); err != nil {
		Trace(err, query, id)
	}
}

// TestTableCollege 测试College表。
func (engine *Engine) TestTableCollege() {
	log.Println("Testing table College.")

	// 准备测试环境。
	const testCollegeName = "（测试学院）"
	engine.DeleteCollegeByName(testCollegeName)

	// 测试CREATE。
	engine.CreateCollege(testCollegeName)
	if !engine.ExistCollege(testCollegeName) {
		log.Panicln("Table College test failed! College with testCollegeName should exist.")
	}

	// 测试READ。
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

	// 测试Delete。
	engine.DeleteCollegeByName(testCollegeName)
	if engine.ExistCollege(testCollegeName) {
		log.Panicln("Table College test failed! College with name testCollegeName should NOT exist.")
	}
}
