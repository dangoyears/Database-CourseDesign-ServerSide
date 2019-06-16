package dbcd

import (
	"log"
	"time"
)

// StudentInfo 学生信息
type StudentInfo struct {
	HumanID        int
	ClassID        int
	StudentNumber  int    `json:"studentId"`
	CollegeName    string `json:"college"`
	SpecialtyName  string `json:"specialty"`
	Grade          int    `json:"grade"`
	ClassCode      int    `json:"class"`
	Name           string `json:"name"`
	Sex            string `json:"sex"`
	Birthday       time.Time
	AdmissionDate  time.Time
	GraduationDate time.Time
	StudentDegree  string
	Status         string
	Notes          string
	PasswordHash   string
	YearOfSchool   int
}

// GetStudentInfo 返回学生信息。
func (engine *Engine) GetStudentInfo() []StudentInfo {
	query := `select "StudentNumber", "CollegeName", "SpecialtyName", "Grade", "ClassCode", "Name", "Sex", "Birthday" from "StudentInfo"`
	rows, err := engine.db.Query(query)
	if err != nil {
		log.Println(query, err)
	}
	defer rows.Close()

	var studentInfo []StudentInfo
	for rows.Next() {
		var info StudentInfo
		if err := rows.Scan(&info.StudentNumber, &info.CollegeName, &info.SpecialtyName, &info.Grade, &info.ClassCode, &info.Name, &info.Sex, &info.Birthday); err != nil {
			log.Println(query, err)
		}
		studentInfo = append(studentInfo, info)
	}

	return studentInfo
}
