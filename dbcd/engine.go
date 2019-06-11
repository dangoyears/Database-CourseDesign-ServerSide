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
	engine.testDB()
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
}

// testDB 测试数据库功能，包括测试数据库能否正常写入等。
func (engine *Engine) testDB() {
	log.Println("Database tests began.")
	TestInsertIntoAcademicYear(engine.db)
	log.Println("All database tests finished.")
}

// establishRouter 建立路由。
func (engine *Engine) establishRouter() {
	engine.router = gin.New()
	engine.router.Use(gin.Logger())
	engine.router.Use(gin.Recovery())

	// CORS
	engine.router.Use(cors.Default())

	engine.BindRoute("/", []string{}, engine.GetWelcomeEndpoint())
	engine.BindRoute("/welcome", []string{}, engine.GetWelcomeEndpoint())
	engine.BindRoute("/login", []string{}, engine.GetLoginEndpoint())
	engine.BindRoute("/logout", []string{}, engine.GetLogoutEndpoint())
	engine.BindRoute("/role", []string{}, engine.GetRoleEndpoint())
	engine.BindRoute("/admin", []string{"admin"}, engine.GetAdminEndpoint())
}

func (engine *Engine) establishGateKeeper() {
	engine.keeper = NewGateKeeper()
}

// Run 启动数据处理引擎。
func (engine *Engine) Run() {
	engine.router.Run("localhost:12323")
}
