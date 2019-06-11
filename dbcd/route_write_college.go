package dbcd

import (
	"github.com/gin-gonic/gin"
)

// GetWriteCollegeEndpoint 返回“/write/college”处的路由。
// @未完成
func (engine *Engine) GetWriteCollegeEndpoint() gin.HandlerFunc {

	type writeCollegeEndpointParam struct {
		College   string `form:"token" binding:"required"`
		Specialty string
		Grade     string
		Class     string
	}

	return func(c *gin.Context) {

	}
}
