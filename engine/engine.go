package engine

import (
	"database/sql"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// goracle.v2 Oracle数据库驱动
	_ "gopkg.in/goracle.v2"

	"github.com/dangoyears/Database-CourseDesign-ServerSide/model"
)

// Engine 数据处理引擎
type Engine struct {
	config Configuration
	db     *sql.DB
	router *gin.Engine
}

// NewEngine 建立数据处理引擎
func NewEngine(config Configuration) Engine {
	var engine Engine
	engine.config = config
	return engine
}

// Start 启动数据处理引擎
func (engine *Engine) Start() {
	engine.establishDB()
	engine.testDB()
	engine.establishRouter()
	engine.router.Run("localhost:12323")
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

func (engine *Engine) establishRouter() {
	engine.router = gin.New()
	engine.router.Use(gin.Logger())
	engine.router.Use(gin.Recovery())

	// CORS
	engine.router.Use(cors.Default())

	engine.router.Run("localhost:12323")
}

func (engine Engine) testDB() {
	model.TestInsertIntoAcademicYear(engine.db)
}
