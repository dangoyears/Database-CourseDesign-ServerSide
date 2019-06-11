package dbcd

import (
	"github.com/gin-gonic/gin"
)

// BindRoute 为路径绑定路由，尚未实现访问权限检查。
func (engine *Engine) BindRoute(path string, grantedLoginType []string, route func(*gin.Context)) {	
	engine.router.GET(path, route)
	engine.router.POST(path, route)
}


// checkAccessPermission 检查当前登陆类型是否有操作对应路由的权限。
func checkAccessPermission(c *gin.Context) {

}
