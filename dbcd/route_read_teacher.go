package dbcd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetReadTeacherEndpoint 返回“/read/teacher”处的路由。
// @未完成
func (engine *Engine) GetReadTeacherEndpoint() gin.HandlerFunc {

	type writeCollegeEndpointParam struct {
		College   string `form:"token" binding:"required"`
		Specialty string
		Grade     string
		Class     string
	}

	var fake = []gin.H{
		{
			"college":    "计算机科学与网络工程学院",
			"name":       "xxx",
			"jobId":      "0000000001",
			"sex":        "男",
			"education":  "硕士",
			"graduation": "南开大学",
			"birthday":   "xxxx-xx-xx",
			"age":        "xx",
			"idCard":     "440582199708310612",
		},
		{
			"college":    "人文学院",
			"name":       "xxx",
			"jobId":      "0000000002",
			"sex":        "女",
			"education":  "博士后",
			"graduation": "北京大学",
			"birthday":   "xxxx-xx-xx",
			"age":        "xx",
			"idCard":     "440582199708310612",
		},
	}

	return func(c *gin.Context) {
		var response = NewRouterResponse()
		response["data"] = fake
		response.SetCodeAndMsg(0, "返回假数据。")
		c.JSON(http.StatusAccepted, response)
	}
}
