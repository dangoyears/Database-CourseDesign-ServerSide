package dbcd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetReadStudentEndpoint 返回“/read/student”处的路由。
// @未完成
func (engine *Engine) GetReadStudentEndpoint() gin.HandlerFunc {

	type writeCollegeEndpointParam struct {
		College   string `form:"token" binding:"required"`
		Specialty string
		Grade     string
		Class     string
	}

	var fake = []gin.H{
		{
			"college":   "计算机科学与网络工程学院",
			"specialty": "软件工程",
			"grade":     "17",
			"class":     "1",
			"name":      "xxx",
			"studentId": "1706300032",
			"status":    "在读本科",
			"sex":       "男",
			"birthday":  "xxxx-xx-xx",
			"age":       "21",
			"idCard":    "440582199708310612",
		},
		{
			"college":   "人文学院",
			"specialty": "汉语言文学",
			"grade":     "18",
			"class":     "1",
			"name":      "xxx",
			"studentId": "1806300027",
			"status":    "在读本科",
			"sex":       "女",
			"birthday":  "xxxx-xx-xx",
			"age":       "21",
			"idCard":    "440582199708310612",
		},
	}

	return func(c *gin.Context) {
		var response = NewRouterResponse()
		response["data"] = fake
		response.SetCodeAndMsg(0, "返回假数据。")
		c.JSON(http.StatusAccepted, response)
	}
}
