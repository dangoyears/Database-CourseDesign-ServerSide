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

// NewEngine 建立数据处理引擎
func NewEngine(config EngineConfiguration) *Engine {
	var engine Engine
	engine.setupConfiguration(config)
	engine.establishDB()
	engine.testDB()
	engine.establishRouter()
	engine.establishGateKeeper()
	return &engine
}

func (engine *Engine) setupConfiguration(config EngineConfiguration) {
	engine.config = config
}

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

func (engine *Engine) testDB() {
	log.Println("Database tests began.")
	TestInsertIntoAcademicYear(engine.db)
	log.Println("All database tests finished.")
}

func (engine *Engine) establishRouter() {
	engine.router = gin.New()
	engine.router.Use(gin.Logger())
	engine.router.Use(gin.Recovery())

	// CORS
	engine.router.Use(cors.Default())

	engine.BindRoute("/login", []string{}, engine.getLoginEndpoint())
	engine.BindRoute("/status/login", []string{}, engine.getLoginStatusEndpoint())
	engine.BindRoute("/logout", []string{}, engine.getLogoutEndpoint())
}

func (engine *Engine) establishGateKeeper() {
	engine.keeper = NewGateKeeper()
}

// Run 启动数据处理引擎。
func (engine *Engine) Run() {
	engine.router.Run("localhost:12323")
}