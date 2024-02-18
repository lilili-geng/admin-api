package role

import (
	"LiadminApi/routes"
	"log"

	"github.com/gin-gonic/gin"
)

// 注册RouterRole
func init() {
	log.Println("init role router success 🚀 ")
	routes.Register(&RouterRole{})
}

type RouterRole struct {
}

func (*RouterRole) Route(r *gin.Engine) {
	h := &HandlerRole{}

	authUserRouter := r.Group("/admin")

	// 角色list
	authUserRouter.GET("/getRoleList", h.getRoleList)

	// 修改
	authUserRouter.POST("/updateRole", h.updateRole)

	// 新增
	
}

// Module 方法返回用户模块标识
func (*RouterRole) Module() string {
	return "role"
}
