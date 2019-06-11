package dbcd

import (
	"strings"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BindRoute 为路径绑定路由。
// path 指定路径。
// grantedRole 指定有哪些角色可以访问路由，留空则允许任何角色访问。
// route 指定路由。
func (engine *Engine) BindRoute(path string, grantedRoles []string, route func(*gin.Context)) {	
	engine.router.GET(path, engine.getPermissionCheckRoute(grantedRoles), route)
	engine.router.POST(path, engine.getPermissionCheckRoute(grantedRoles), route)
}

// getPermissionCheckRoute 检查
func (engine *Engine) getPermissionCheckRoute(grantedRoles []string)  gin.HandlerFunc {
	type permissionParam struct {
		Token string `form:"token"`
	}
	
	return func(c *gin.Context) {	
		if len(grantedRoles) == 0 {
			return  // 允许任何人访问，将控制权交给下一个路由。
		}
		
		var param permissionParam
		if c.ShouldBind(&param) == nil {
			comeInRole := engine.keeper.GetRole(param.Token)
			for _, grantedRole := range grantedRoles {
				if comeInRole == grantedRole {
					return  // 允许访问，将控制权交给下一个路由。 
				}
			}
		}
		
		// 拒绝访问
		var response = NewRouterResponse()
		response.SetCodeAndMsg(-1, "权限不足。只允许" +strings.Join(grantedRoles, "、") + "角色访问。")
		c.JSON(http.StatusOK, response)
		c.Abort()
	}
}
