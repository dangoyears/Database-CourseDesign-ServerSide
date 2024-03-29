package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"runtime/debug"

	"github.com/dangoyears/Database-CourseDesign-ServerSide/dbcd"
	_ "gopkg.in/goracle.v2"
	"gopkg.in/yaml.v2"
)

// main 是应用程序的入口。
// 从配置文件config.yaml中加载配置并验证，
// 随后启动数据处理引擎。
func main() {
	config := loadConfiguration()
	verifyConfiguration(config)
	engine := dbcd.NewEngine(config)
	engine.Run()
}

// loadConfiguration 从config.yaml中加载配置信息。
func loadConfiguration() (config dbcd.EngineConfiguration) {
	log.Println("Loading configuration from file: config.yaml")
	configYaml, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalln(err)
		debug.PrintStack()
	}
	err = yaml.Unmarshal(configYaml, &config)
	if err != nil {
		log.Fatalln(err)
		debug.PrintStack()
	}
	return config
}

// verifyConfiguration 验证当前加载的配置信息是否正确。
func verifyConfiguration(config dbcd.EngineConfiguration) {
	log.Println("Verifying configuration...")
	db, err := sql.Open("goracle", config.OracleConnectString)
	defer db.Close()

	if err != nil {
		log.Fatalln(err)
		debug.PrintStack()
	}
	if err = db.Ping(); err != nil {
		log.Fatalln(err)
		debug.PrintStack()
	}
	log.Println("Configuration is verified.")
}
