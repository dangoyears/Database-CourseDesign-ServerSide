package dbcd

import "log"

// ClassInfo 是视图“ClassInfo”的抽象
type ClassInfo struct {
	CollegeName       string
	SpecialtyName     string
	Grade             int
	ClassCode         int
	TotalStudentCount int
}

// GetClassInfo 返回班级信息。
func (engine *Engine) GetClassInfo() []ClassInfo {
	query := `select "CollegeName", "SpecialtyName", "Grade", "ClassCode", "TotalStudentCount" from "ClassInfo"`
	rows, err := engine.db.Query(query)
	if err != nil {
		log.Println(query, err)
	}
	defer rows.Close()

	var classInfo []ClassInfo
	for rows.Next() {
		var info ClassInfo
		if err := rows.Scan(&info.CollegeName, &info.SpecialtyName, &info.Grade, &info.ClassCode, &info.TotalStudentCount); err != nil {
			log.Println(query, err)
		}
		classInfo = append(classInfo, info)
	}

	return classInfo
}
