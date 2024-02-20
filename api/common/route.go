package common

import (
	"LiadminApi/routes"
	"log"

	"github.com/gin-gonic/gin"
)

// 公共接口
// 注册RouterCommon
func init() {
	log.Println("init common router success 🚀 ")
	routes.Register(&RouterCommon{})
}

type RouterCommon struct {
}

func (*RouterCommon) Route(r *gin.Engine) {
	h := &HandlerCommon{}
	authCommon := r.Group("/api")
	//注册
	authCommon.POST("/register", h.registerUser)
	// 登陆
	authCommon.POST("/login", h.login)
}

func (*RouterCommon) Module() string {
	return "common" // 返回模块标识
}
