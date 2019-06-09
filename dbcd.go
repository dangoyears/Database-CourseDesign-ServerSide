package main

import (
	"database/sql"
	"io/ioutil"
	"log"

	"github.com/dangoyears/Database-CourseDesign-ServerSide/engine"
	_ "gopkg.in/goracle.v2"
	"gopkg.in/yaml.v2"
)

func main() {
	config := loadConfiguration()
	verifyConfiguration(config)
	engine := engine.NewEngine(config)
	engine.Run()
}

func loadConfiguration() (config engine.Configuration) {
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

func verifyConfiguration(config engine.Configuration) {
	db, err := sql.Open("goracle", config.OracleConnectString)
	defer db.Close()

	if err != nil {
		log.Fatalln(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalln(err)
	}
}
