package dbcd

import (
	"net/http"
	"strings"

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

// getPermissionCheckRoute 检查一个API请求是否符合访问权限管理的要求。
func (engine *Engine) getPermissionCheckRoute(grantedRoles []string) gin.HandlerFunc {
	type permissionParam struct {
		Token string `json:"token" form:"token" binding:"required"`
	}

	return func(c *gin.Context) {
		// 需要改进
		// https://gin-gonic.com/docs/examples/bind-body-into-dirrerent-structs/
		//

		if len(grantedRoles) == 0 { // 公开API
			return // 允许任何人访问，将控制权交给下一个路由。
		}

		var param permissionParam
		if c.ShouldBind(&param) == nil {
			comeInRole := engine.keeper.GetRole(param.Token)
			for _, grantedRole := range grantedRoles {
				if comeInRole == grantedRole {
					c.Next()
					return // 允许访问，将控制权交给下一个路由。
				}
			}

			// 拒绝访问。
			var response = NewRouterResponse()
			response.SetCodeAndMsg(-1, "权限不足。只允许"+strings.Join(grantedRoles, "、")+"角色访问。传入token的权限不足。")
			c.JSON(http.StatusOK, response)
			c.Abort()
			return
		}

		// 未传入token参数，拒绝访问。
		var response = NewRouterResponse()
		response.SetCodeAndMsg(-1, "权限不足。只允许"+strings.Join(grantedRoles, "、")+"角色访问。请传入token以验证用户身份。")
		c.JSON(http.StatusOK, response)
		c.Abort()
	}
}
