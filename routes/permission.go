package routes

import (
	"LiadminApi/middleware"
	"LiadminApi/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)
func PermissionMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求的path
		p := c.Request.URL.Path
		// 请求的方法
		m := c.Request.Method

		name, _ := c.Get("name")

		fmt.Println(name,"namenamename")

		role:="superAdmin"
		// role:="user"
		// role := "guest"

		fmt.Println("role:" + role)
		fmt.Println("path:" + p)
		fmt.Println("method:" + m)

		rsp := &utils.Result{}

		// 检查用户权限
		isPass := middleware.Enforcer.Enforce(role, p, m)
		if !isPass {
			c.JSON(401, rsp.Fail(401, "无访问权限"))
			c.Abort() // 中止请求链
			return
		}

		// 权限检查通过，继续执行下一个中间件或处理程序
		c.Next()
	}
}

