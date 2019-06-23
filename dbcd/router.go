package dbcd

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func saveRequestBody(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		log.Panicln(err)
	}
	c.Set("request-body", body)
	log.Println("Finish setting up body.")
}

func resumeRequestBody(c *gin.Context) {
	body, exists := c.Get("request-body")

	if !exists {
		log.Panicln("Request body should exist. Call saveRequestBody() before resumRequestBody().")
	}
	bodyBytes := body.([]byte)

	c.Request.Body = ioutil.NopCloser(bytes.NewReader(bodyBytes))
}

// BindContextIntoStruct 尝试将Context与Struct绑定。
func BindContextIntoStruct(c *gin.Context, obj interface{}) *error {
	var err1, err2 error
	if err1 = c.ShouldBindWith(obj, binding.Query); err1 == nil {
		Trace(obj)
		return nil
	} else if err2 = c.ShouldBindBodyWith(obj, binding.JSON); err2 == nil {
		Trace(obj)
		return nil
	}
	var err = errors.New(strings.Join([]string{err1.Error(), err2.Error()}, " "))
	return &err
}

// GetArgs 从Context中getArgsParseRoute设定的网页参数。
func GetArgs(c *gin.Context) map[string]interface{} {
	param, exists := c.Get("args")
	if !exists {
		log.Panicln("Args should always exist.")
	}

	return param.(map[string]interface{})
}

// BindRoute 为路径绑定路由。
// path 指定路径。
// grantedRole 指定有哪些角色可以访问路由，留空则允许任何角色访问。
// route 指定路由。
func (engine *Engine) BindRoute(path string, grantedRoles []string, route func(*gin.Context)) {
	engine.router.GET(path, engine.getPermissionCheckRoute(grantedRoles), route)
	engine.router.POST(path, engine.getPermissionCheckRoute(grantedRoles), route)
}

// getJSONParseRoute 返回一个参数绑定路由，这个路由从查询字符串和JSON格式的请求主体中获取网页参数，
// 并将参数保存到Context的“Args”变量中。
func (engine *Engine) getArgsParseRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		args := make(map[string]interface{})

		// 解析来自查询字符串的参数。
		for key, value := range c.Request.URL.Query() {
			args[key] = value[0] // 传入多个同名参数的情况下取第一个参数。
		}

		// 解析来自JSON的参数。
		buf := new(bytes.Buffer)
		buf.ReadFrom(c.Request.Body)
		if err := json.Unmarshal(buf.Bytes(), &args); err != nil {
			Trace(err)
		}

		// 尝试将所有可以转换为数值类型的参数转换为数值类型。
		for key, value := range args {
			str, ok := value.(string)
			if ok {
				intval, err := strconv.Atoi(str)
				if err == nil {
					args[key] = intval
				}
			}
		}

		c.Set("args", args)
	}
}

// getPermissionCheckRoute 检查一个API请求是否符合访问权限管理的要求。
func (engine *Engine) getPermissionCheckRoute(grantedRoles []string) gin.HandlerFunc {
	type permissionParam struct {
		Token string `json:"token" form:"token" binding:"required"`
	}

	return func(c *gin.Context) {
		saveRequestBody(c)
		resumeRequestBody(c)

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
