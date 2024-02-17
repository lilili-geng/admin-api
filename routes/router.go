package routes

import (
	"LiadminApi/docs"
	"LiadminApi/middleware"
	"LiadminApi/utils"
	"fmt"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Router接口
type Router interface {
	Route(r *gin.Engine)
	Module() string // 模块标识
}

// 传入
type RouterRegister struct{}

func New() *RouterRegister {
	return &RouterRegister{}
}

func (*RouterRegister) Route(ro Router, r *gin.Engine) {
	ro.Route(r)
}

// 公共路由注册
var routers []Router

func InitRouter(r *gin.Engine) {

	// 中间件
	r.Use(middleware.Cors()).Use(Recover).Use(AccessLog())

	for _, ro := range routers {
		fmt.Println("routersrouters", ro.Module())
		if ro.Module() != "common" {
			r.Use(utils.JwtToken()).Use(PermissionMiddleWare())
		}
		ro.Route(r)
	}

	// 找不到路由
	r.NoRoute(HandleNotFound)
	r.NoMethod(HandleNotFound)

	// swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func Register(ro ...Router) {
	routers = append(routers, ro...)
}

// 404
func HandleNotFound(ctx *gin.Context) {
	middleware.Logger.Errorf("handle not found: %v", ctx.Request.RequestURI)
	middleware.Logger.Errorf("stack: %v", string(debug.Stack()))
	rsp := &utils.Result{}
	ctx.JSON(200, rsp.Fail(404, "资源未找到"))
	return
}

// 500
func Recover(ctx *gin.Context) {
	defer func() {
		rsp := &utils.Result{}
		if r := recover(); r != nil {
			//打印错误堆栈信息
			//log.Printf("panic: %v\n", r)
			middleware.Logger.Errorf("panic: %v", r)
			//log stack
			middleware.Logger.Errorf("stack: %v", string(debug.Stack()))
			//print stack
			debug.PrintStack()
			//return
			ctx.JSON(200, rsp.Fail(500, "服务器内部错误"))
		}
	}()
	//继续后续接口调用
	ctx.Next()
}
