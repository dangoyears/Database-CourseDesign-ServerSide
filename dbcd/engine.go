package dbcd

import (
	"database/sql"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// goracle.v2 Oracle数据库驱动
	_ "gopkg.in/goracle.v2"
)

// Engine 数据处理引擎
type Engine struct {
	config EngineConfiguration
	db     *sql.DB
	router *gin.Engine
	keeper *GateKeeper
}

// EngineConfiguration 数据处理引擎配置
type EngineConfiguration struct {
	OracleConnectString string `yaml:"OracleConnectString"`
}

// NewEngine 返回数据处理引擎实例。
func NewEngine(config EngineConfiguration) *Engine {
	var engine Engine
	engine.setupConfiguration(config)
	engine.establishDB()
	engine.ensureThatAtLeastOneAdministratorExists()
	engine.establishRouter()
	engine.establishGateKeeper()
	return &engine
}

// setupConfiguration 加载配置文件。
func (engine *Engine) setupConfiguration(config EngineConfiguration) {
	engine.config = config
}

// establishDB 建立数据库连接。
func (engine *Engine) establishDB() {
	var err error
	engine.db, err = sql.Open("goracle", engine.config.OracleConnectString)
	if err != nil {
		log.Fatalln(err)
	}
	if err = engine.db.Ping(); err != nil {
		log.Fatalln(err)
	}
	engine.testDB()
}

// testDB 测试数据库功能，包括测试数据库能否正常写入等。
func (engine *Engine) testDB() {
	log.Println("Database tests began.")
	engine.TestTableAdministrator()
	engine.TestTableClass()
	engine.TestTableCollege()
	engine.TestTableCourse()
	engine.TestTableHuman()
	engine.TestTableSpecialty()
	engine.TestTableStudent()
	engine.TestTableTeacher()
	engine.TestViewClassInfo()
	engine.TestViewStudentInfo()
	engine.TestViewTeacherInfo()
	log.Println("All database tests have finished.")
}

// ensureThatAtLeastOneAdministratorExists 确保至少有一个管理员账户存在。
// 如果不存在任意的管理员账户，则创建默认账户。
func (engine *Engine) ensureThatAtLeastOneAdministratorExists() {
	if engine.GetAdministratorsCount() == 0 {
		log.Println("Create default administrator account since there is no admin exists.")
		engine.CreateDefaultAdministrator()
	}
}

// establishRouter 建立路由。
func (engine *Engine) establishRouter() {
	gin.SetMode(gin.ReleaseMode)
	engine.router = gin.New()
	engine.router.Use(gin.Logger())
	engine.router.Use(gin.Recovery())

	// CORS
	engine.router.Use(cors.Default())

	engine.BindRoute("/", []string{}, engine.GetWelcomeEndpoint())
	engine.BindRoute("/welcome", []string{}, engine.GetWelcomeEndpoint())
	engine.BindRoute("/echo", []string{"admin"}, engine.getEchoRoute())

	engine.BindRoute("/login", []string{}, engine.GetLoginEndpoint())
	engine.BindRoute("/logout", []string{}, engine.GetLogoutEndpoint())
	engine.BindRoute("/role", []string{}, engine.GetRoleEndpoint())
	engine.BindRoute("/admin", []string{"admin"}, engine.GetAdminEndpoint())

	engine.BindRoute("/read/college", []string{"admin"}, engine.GetReadCollegeEndpoint())
	engine.BindRoute("/read/student", []string{"admin"}, engine.GetReadStudentEndpoint())
	engine.BindRoute("/read/teacher", []string{"admin"}, engine.GetReadTeacherEndpoint())
	engine.BindRoute("/read/course", []string{"admin"}, engine.GetReadCourseEndpoint())

	engine.BindRoute("/read/teacher/one", []string{"admin", "teacher"}, engine.GetReadTeacherOneEndpoint())
	engine.BindRoute("/read/student/one", []string{"admin", "student"}, engine.GetReadStudentOneEndpoint())

	engine.BindRoute("/write/college", []string{"admin"}, engine.GetWriteCollegeEndpoint())
	engine.BindRoute("/write/teacher", []string{"admin"}, engine.GetWriteTeacherEndpoint())
	engine.BindRoute("/write/student", []string{"admin"}, engine.GetWriteStudentEndpoint())
	engine.BindRoute("/write/course", []string{"admin", "teacher"}, engine.GetWriteCourseEndpoint())

	engine.BindRoute("/delete/class", []string{"admin"}, engine.GetDeleteClassEndpoint())
	engine.BindRoute("/delete/both", []string{"admin"}, engine.GetDeleteBothEndpoint())
}

// establishGateKeeper 建立用户状态管理器。
func (engine *Engine) establishGateKeeper() {
	engine.keeper = NewGateKeeper()
}

// Run 启动数据处理引擎。
func (engine *Engine) Run() {
	log.Println("Engine is online.")
	engine.router.Run("localhost:12323")
}
