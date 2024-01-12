package routes

import (

	"github.com/gin-gonic/gin"
)

// Router接口
type Router interface {
	Route(r *gin.Engine)
}

// 传入
type RouterRegister struct {
}

func New() *RouterRegister {
	return &RouterRegister{}
}

func (*RouterRegister) Route(ro Router, r *gin.Engine) {
	ro.Route(r)
}

// 公共路由注册
var routers []Router

func InitRouter(r *gin.Engine) {
	for _, ro := range routers {
		ro.Route(r)
	}
}

func Register(ro ...Router) {
	routers = append(routers, ro...)
}
