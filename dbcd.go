package main

import (
	"database/sql"
	"io/ioutil"
	"log"

	"github.com/dangoyears/Database-CourseDesign-ServerSide/dbcd"
	_ "gopkg.in/goracle.v2"
	"gopkg.in/yaml.v2"
)

// 从config.yaml中加载配置文件并验证内容，
// 随后启动数据处理引擎
func main() {
	config := loadConfiguration()
	verifyConfiguration(config)
	engine := dbcd.NewEngine(config)
	engine.Run()
}

func loadConfiguration() (config dbcd.EngineConfiguration) {
	configYaml, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(configYaml, &config)
	if err != nil {
		log.Fatalln(err)
	}
	return config
}

func verifyConfiguration(config dbcd.EngineConfiguration) {
	db, err := sql.Open("goracle", config.OracleConnectString)
	defer db.Close()

	if err != nil {
		log.Fatalln(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalln(err)
	}
}
