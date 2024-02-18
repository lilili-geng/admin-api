package routes

import (
	"LiadminApi/middleware"
	"LiadminApi/service"
	"LiadminApi/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 权限验证
func PermissionMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求的path
		p := c.Request.URL.Path
		// 请求的方法
		m := c.Request.Method

		nameInterface, _ := c.Get("name")

		name, ok := nameInterface.(string)

		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to convert username to string"})
			return
		}

		user, _ := service.GetUserInfo(name)

		// role := "superAdmin"
		// role:="user"
		// role := "guest"

		// fmt.Println("role:" + user.RoleName)
		// fmt.Println("path:" + p)
		// fmt.Println("method:" + m)

		rsp := &utils.Result{}

		if user == nil {
			c.JSON(401, rsp.Fail(401, "无访问权限"))
			c.Abort() // 中止请求链
			return
		}


		fmt.Println("role user.RoleName :",user.RoleName)
		fmt.Println(user)

		// 检查用户权限
		isPass := middleware.Enforcer.Enforce(user.RoleName, p, m)

		if !isPass {
			c.JSON(401, rsp.Fail(401, "无访问权限"))
			c.Abort() // 中止请求链
			return
		}

		// 权限检查通过，继续执行下一个中间件或处理程序
		c.Next()
	}
}
