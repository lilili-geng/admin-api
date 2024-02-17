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
	// 角色list
	r.GET("/getByUserList", h.getByUserList)
}

// Module 方法返回用户模块标识
func (*RouterRole) Module() string {
	return "role"
}
