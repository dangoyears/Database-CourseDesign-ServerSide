package dbcd

import (
	"log"
	"time"
)

// TeacherInfo 教师信息
type TeacherInfo struct {
	HumanID       int
	CollegeID     int
	TeacherNumber int    `json:"jobId"`
	CollegeName   string `json:"college"`
	Name          string `json:"name"`
	Sex           string `json:"sex"`
	Birthday      time.Time
	Identity      string
	Notes         string
	PasswordHash  string
}

// GetTeacherInfo 返回教师信息。
func (engine *Engine) GetTeacherInfo() []TeacherInfo {
	query := `select "TeacherNumber", "CollegeName", "Name", "Sex", "Birthday", "Identity" from "TeacherInfo"`
	rows, err := engine.db.Query(query)
	if err != nil {
		log.Println(query, err)
	}
	defer rows.Close()

	var teacherInfo []TeacherInfo
	for rows.Next() {
		var info TeacherInfo
		if err := rows.Scan(&info.TeacherNumber, &info.CollegeName, &info.Name, &info.Sex, &info.Birthday, &info.Identity); err != nil {
			log.Println(query, err)
		}
		teacherInfo = append(teacherInfo, info)
	}

	return teacherInfo
}
