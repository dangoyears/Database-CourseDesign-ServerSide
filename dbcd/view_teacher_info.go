package dbcd

import (
	"log"
	"time"
)

// TeacherInfo 教师信息
type TeacherInfo struct {
	HumanID          int
	CollegeID        int
	CollegeName      string `json:"college"`
	Name             string `json:"name"`
	Sex              string `json:"sex"`
	Birthday         time.Time
	Identity         string
	Notes            string
	PasswordHash     string
	TeacherNumber    int `json:"jobId"`
	GraduationSchool string
	Position         string
	TeacherDegree    string
}

// CreateTeacherAsInfo 根据教师信息创建教师
func (engine *Engine) CreateTeacherAsInfo(info TeacherInfo) {
	var human = Human{
		Name:         info.Name,
		Sex:          info.Sex,
		Birthday:     info.Birthday,
		Identity:     info.Identity,
		Notes:        info.Notes,
		PasswordHash: info.PasswordHash,
	}

	engine.CreateTeacher(human, info.CollegeName, info.TeacherNumber, info.GraduationSchool, info.Position, info.TeacherDegree)
}

// GetTeacherInfo 返回教师信息。
func (engine *Engine) GetTeacherInfo() []TeacherInfo {
	query := `select "HumanID", "CollegeID", "CollegeName", "Name", "Sex", "Birthday", "Identity", "Notes", "PasswordHash", 
"TeacherNumber", "GraduationSchool", "Position", "TeacherDegree" from "TeacherInfo"`
	rows, err := engine.db.Query(query)
	if err != nil {
		log.Println(query, err)
	}
	defer rows.Close()

	var teacherInfo []TeacherInfo
	for rows.Next() {
		var info TeacherInfo
		if err := rows.Scan(
			&info.HumanID, &info.CollegeID,
			&info.CollegeName, &info.Name, &info.Sex, &info.Birthday, &info.Identity, &info.Notes, &info.PasswordHash,
			&info.TeacherNumber, &info.GraduationSchool, &info.Position, &info.TeacherDegree); err != nil {

			Trace(query, err)
		}
		teacherInfo = append(teacherInfo, info)
	}

	return teacherInfo
}

// GetTeacherInfoByTeacherNumber 返回对应学号的教师信息。
func (engine *Engine) GetTeacherInfoByTeacherNumber(teacherNumber int) *TeacherInfo {
	query := `select "HumanID", "CollegeID", "CollegeName", "Name", "Sex", "Birthday", "Identity", "Notes", "PasswordHash", 
"TeacherNumber", "GraduationSchool", "Position", "TeacherDegree" from "TeacherInfo" where "TeacherNumber"=:1`
	row := engine.db.QueryRow(query, teacherNumber)

	var info TeacherInfo
	if err := row.Scan(&info.HumanID, &info.CollegeID,
		&info.CollegeName, &info.Name, &info.Sex, &info.Birthday, &info.Identity, &info.Notes, &info.PasswordHash,
		&info.TeacherNumber, &info.GraduationSchool, &info.Position, &info.TeacherDegree); err != nil {

		Trace(query, err)
		return nil
	}

	return &info
}

// GetTeacherInfoByTeacherName 返回指定teacherName的教师。这里假设教师没有重名的情况。
func (engine *Engine) GetTeacherInfoByTeacherName(teacherName string) *TeacherInfo {
	query := `select "HumanID", "CollegeID", "CollegeName", "Name", "Sex", "Birthday", "Identity", "Notes", "PasswordHash", 
"TeacherNumber", "GraduationSchool", "Position", "TeacherDegree" from "TeacherInfo" where "Name"=:1`
	row := engine.db.QueryRow(query, teacherName)

	var info TeacherInfo
	if err := row.Scan(&info.HumanID, &info.CollegeID,
		&info.CollegeName, &info.Name, &info.Sex, &info.Birthday, &info.Identity, &info.Notes, &info.PasswordHash,
		&info.TeacherNumber, &info.GraduationSchool, &info.Position, &info.TeacherDegree); err != nil {

		Trace(query, err)
		return nil
	}

	return &info
}

// UpdateTeacherAsInfo 更新指定教师的信息。
// 可以处理教师转院等情况。
// 禁止修改教师工号。
func (engine *Engine) UpdateTeacherAsInfo(teacherNumber int, info TeacherInfo) {
	var human = Human{
		Name:         info.Name,
		Sex:          info.Sex,
		Birthday:     info.Birthday,
		Identity:     info.Identity,
		Notes:        info.Notes,
		PasswordHash: info.PasswordHash,
	}

	engine.UpdateTeacher(teacherNumber, human, info.CollegeName, info.GraduationSchool, info.Position, info.TeacherDegree)
}

// TestViewTeacherInfo 测试TeacherInfo视图。
func (engine *Engine) TestViewTeacherInfo() {
	log.Println("Testing view TeacherInfo.")

	engine.GetTeacherInfo()
}
