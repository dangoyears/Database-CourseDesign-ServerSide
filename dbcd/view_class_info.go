package dbcd

import "log"

// ClassInfo 是视图“ClassInfo”的抽象
type ClassInfo struct {
	CollegeID         int
	SpecialtyID       int
	ClassID           int
	CollegeName       string `json:"college"`
	SpecialtyName     string `json:"specialty"`
	Grade             int    `json:"grade"`
	ClassCode         int    `json:"class"`
	TotalStudentCount int    `json:"sum"`
}

// GetClassInfo 返回班级信息。
func (engine *Engine) GetClassInfo() []ClassInfo {
	query := `select "CollegeID", "SpecialtyID", "ClassID", "CollegeName", "SpecialtyName", "Grade", "ClassCode", "TotalStudentCount" from "ClassInfo"`
	rows, err := engine.db.Query(query)
	if err != nil {
		log.Println(query, err)
	}
	defer rows.Close()

	var classInfo []ClassInfo
	for rows.Next() {
		var info ClassInfo
		if err := rows.Scan(&info.CollegeID, &info.SpecialtyID, &info.ClassID,
			&info.CollegeName, &info.SpecialtyName, &info.Grade, &info.ClassCode, &info.TotalStudentCount); err != nil {
			log.Println(query, err)
		}
		classInfo = append(classInfo, info)
	}

	return classInfo
}

// TestViewClassInfo 测试Class视图。
func (engine *Engine) TestViewClassInfo() {
	log.Println("Testing view ClassInfo.")

	engine.GetClassInfo()
}
