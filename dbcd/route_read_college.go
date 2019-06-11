package dbcd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetReadCollegeEndpoint 返回“/read/college”处的路由。
// @未完成
func (engine *Engine) GetReadCollegeEndpoint() gin.HandlerFunc {

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
			"sum":       "41",
		},
		{
			"college":   "计算机科学与网络工程学院",
			"specialty": "软件工程",
			"grade":     "18",
			"class":     "4",
			"sum":       "40",
		},
		{
			"college":   "人文学院",
			"specialty": "汉语言文学",
			"grade":     "15",
			"class":     "2",
			"sum":       "41",
		},
	}

	return func(c *gin.Context) {
		var response = NewRouterResponse()
		response["data"] = fake
		response.SetCodeAndMsg(0, "返回假数据。")
		c.JSON(http.StatusAccepted, response)
	}
}
