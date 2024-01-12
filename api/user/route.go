package user

import (
	"LiadminApi/routes"
	"log"

	"github.com/gin-gonic/gin"
)

// 注册RouterUser
func init() {
	log.Println("init user router success 🚀 ")
	routes.Register(&RouterUser{})
}

type RouterUser struct {
}

func (*RouterUser) Route(r *gin.Engine) {
	h := &HandlerUser{}
	r.POST("/login", h.getCaptcha)
}
