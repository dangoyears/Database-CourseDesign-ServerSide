package dbcd

import (
	"log"
	"time"
)

// StudentInfo 学生信息
type StudentInfo struct {
	HumanID        int
	CollegeID      int
	SpecialtyID    int
	ClassID        int
	CollegeName    string `json:"college"`
	SpecialtyName  string `json:"specialty"`
	Grade          int    `json:"grade"`
	ClassCode      int    `json:"class"`
	Name           string `json:"name"`
	Sex            string `json:"sex"`
	Birthday       time.Time
	Identity       string
	Notes          string
	PasswordHash   string
	StudentNumber  int `json:"studentId"`
	AdmissionDate  time.Time
	GraduationDate time.Time
	StudentDegree  string
	YearOfSchool   int
	Status         string
}

// CreateStudentAsInfo 根据信息创建学生。
func (engine *Engine) CreateStudentAsInfo(info StudentInfo) {
	var human = Human{
		Name:         info.Name,
		Sex:          info.Sex,
		Birthday:     info.Birthday,
		Identity:     info.Identity,
		Notes:        info.Notes,
		PasswordHash: info.PasswordHash,
	}

	engine.CreateStudent(human, info.CollegeName, info.SpecialtyName, info.Grade, info.ClassCode, info.StudentNumber)
	engine.UpdateStudent(info.StudentNumber, info)
}

// UpdateStudentAsInfo 根据信息更新指定studentNumber的学生。
// 可以处理转专业、转学院等。
// 禁止修改学号。
func (engine *Engine) UpdateStudentAsInfo(studentNumber int, info StudentInfo) {
	var human = Human{
		Name:         info.Name,
		Sex:          info.Sex,
		Birthday:     info.Birthday,
		Identity:     info.Identity,
		Notes:        info.Notes,
		PasswordHash: info.PasswordHash,
	}
	oldStudent := engine.GetStudentInfoByStudentNumber(studentNumber)

	if oldStudent == nil {
		return
	}

	engine.UpdateHumanByID(human, oldStudent.HumanID)
	engine.UpdateStudent(studentNumber, info)

	// 转专业/转学院/转班级
	if !engine.ExistClass(info.CollegeName, info.SpecialtyName, info.Grade, info.ClassCode) {
		engine.CreateClass(info.CollegeName, info.SpecialtyName, info.Grade, info.ClassCode)
	}
	newClassID := engine.GetClassBySpecialtyNameGradeAndCode(info.SpecialtyName, info.Grade, info.ClassCode).ClassID

	queryUpdate := `update "Student" set "ClassID"=:1 where "StudentNumber"=:2`
	if _, err := engine.db.Exec(queryUpdate, newClassID, studentNumber); err != nil {
		Trace(err, queryUpdate, newClassID, studentNumber)
	}
}

// GetStudentInfoByStudentNumber 获取对应学号的学生的信息。
func (engine *Engine) GetStudentInfoByStudentNumber(id int) *StudentInfo {
	query := `select "HumanID", "CollegeID", "SpecialtyID", "ClassID", 
"CollegeName", "SpecialtyName", "Grade", "ClassCode", 
"Name", "Sex", "Birthday", "Identity", "Notes", "PasswordHash",
"StudentNumber", "AdmissionDate", "GraduationDate", "StudentDegree", "YearOfSchool", "Status"
from "StudentInfo"
where "StudentNumber"=:1`
	row := engine.db.QueryRow(query, id)

	var info StudentInfo
	if err := row.Scan(&info.HumanID, &info.CollegeID, &info.SpecialtyID, &info.ClassID,
		&info.CollegeName, &info.SpecialtyName, &info.Grade, &info.ClassCode,
		&info.Name, &info.Sex, &info.Birthday, &info.Identity, &info.Notes, &info.PasswordHash,
		&info.StudentNumber, &info.AdmissionDate, &info.GraduationDate, &info.StudentDegree, &info.YearOfSchool, &info.Status); err != nil {
		Trace(err, query, id)
		return nil
	}
	return &info
}

// GetStudentInfo 返回学生信息。
func (engine *Engine) GetStudentInfo() []StudentInfo {
	query := `select "HumanID", "CollegeID", "SpecialtyID", "ClassID", 
"CollegeName", "SpecialtyName", "Grade", "ClassCode",
"Name", "Sex", "Birthday", "Identity", "Notes", "PasswordHash",
"StudentNumber", "AdmissionDate", "GraduationDate", "StudentDegree", "YearOfSchool", "Status"
from "StudentInfo"`
	rows, err := engine.db.Query(query)
	if err != nil {
		Trace(err, query)
	}
	defer rows.Close()

	var studentInfo []StudentInfo
	for rows.Next() {
		var info StudentInfo
		if err := rows.Scan(
			&info.HumanID, &info.CollegeID, &info.SpecialtyID, &info.ClassID,
			&info.CollegeName, &info.SpecialtyName, &info.Grade, &info.ClassCode,
			&info.Name, &info.Sex, &info.Birthday, &info.Identity, &info.Notes, &info.PasswordHash,
			&info.StudentNumber, &info.AdmissionDate, &info.GraduationDate, &info.StudentDegree, &info.YearOfSchool, &info.Status); err != nil {
			log.Println(query, err)
		}
		studentInfo = append(studentInfo, info)
	}

	return studentInfo
}

// TestViewStudentInfo 测试StudentInfo视图。
func (engine *Engine) TestViewStudentInfo() {
	log.Println("Testing view StudentInfo.")

	engine.GetStudentInfo()
}
