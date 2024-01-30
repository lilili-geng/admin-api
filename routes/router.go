package routes

import (
	"LiadminApi/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// swagger
	docs.SwaggerInfo.BasePath=""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func Register(ro ...Router) {
	routers = append(routers, ro...)
}
