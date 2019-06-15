package dbcd

// Class 是表“Class”的抽象
type Class struct {
	ClassID              int
	SpecialtyID          int
	MasterTeacherHumanID int
	Grade                int
	ClassCode            int
}

// CreateClass 根据专业名specialty、届别grade和班别code来创建班级。
// 返回创建的班级结构。若班级已经存在，则函数返回既有的班级结构。
func (engine *Engine) CreateClass(specialty string, grade, code int)  {
	
}

// getClassByRowID 根据数据库内部行号获取Class。
func (engine *Engine) getClassByRowID(id int) {

}

// TestClass 测试Class表。
func (engine *Engine) TestClass() {

}
