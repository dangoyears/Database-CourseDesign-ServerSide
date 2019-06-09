package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

	"github.com/dangoyears/Database-CourseDesign-ServerSide/model"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "gopkg.in/goracle.v2"
)

type configuation struct {
	OracleConnectString string `yaml:"OracleConnectString"`
}

func loadConfiguration() configuation {
	var config configuation
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

func appendStatus(json *gin.H, statusCode int, statusText string) {
	(*json)["statusCode"] = statusCode
	(*json)["statusText"] = statusText
}

func appendSuccessfulStatus(json *gin.H) {
	appendStatus(json, 0, "成功。这是一个临时路由，数据的写操作不会真正生效(*^_^*)")
}

func appendFailureStatus(json *gin.H) {
	appendStatus(json, -1, "失败。")
}

func main() {
	config := loadConfiguration()

	db, err := sql.Open("goracle", config.OracleConnectString)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select 2+2 from dual")
	if err != nil {
		fmt.Println("Error fetching addition")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var sum int
		rows.Scan(&sum)
		fmt.Printf("2 + 2 always equals: %d\n", sum)
	}

	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/login", func(c *gin.Context) {
		var response = make(gin.H)
		var human model.Human

		if (c.ShouldBind(&human)) == nil && human.Password == "dangoyears" {
			appendSuccessfulStatus(&response)
			response["token"] = "db0c658f6be9d2e94a4efed996779dbe2bc69b27"
		} else {
			appendFailureStatus(&response)
			response["token"] = ""
		}

		c.JSON(200, response)
	})

	r.GET("/write/college", func(c *gin.Context) {
		var response = make(gin.H)
		appendSuccessfulStatus(&response)
		c.JSON(200, response)
	})

	r.POST("/write/college", func(c *gin.Context) {
		var response = make(gin.H)
		appendSuccessfulStatus(&response)
		c.JSON(200, response)
	})

	r.GET("/read/college", func(c *gin.Context) {
		var response = make(gin.H)
		appendSuccessfulStatus(&response)
		var data = []gin.H{
			{
				"college":   "计算机科学与网络工程学院",
				"specialty": "软件工程",
				"grade":     "17级",
				"class":     "1班",
				"sum":       "41",
			},
			{
				"college":   "计算机科学与网络工程学院",
				"specialty": "软件工程",
				"grade":     "18级",
				"class":     "4班",
				"sum":       "40",
			},
			{
				"college":   "人文学院",
				"specialty": "汉语言文学",
				"grade":     "15级",
				"class":     "2班",
				"sum":       "41",
			},
		}
		response["data"] = data

		c.JSON(200, response)
	})

	r.Run("localhost:12323")
}
