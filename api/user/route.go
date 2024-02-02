package user

import (
	"LiadminApi/routes"
	"LiadminApi/utils"
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
	authUserRouter := r.Group("/user")

	authUserRouter.Use(utils.JwtToken())
	
	// 用户list
	authUserRouter.GET("/getByUserList", h.getByUserList)
	// 删除user
	authUserRouter.DELETE("/deleteUserById", h.deleteUserById)
	// 根据id查询用户
	authUserRouter.GET("/getByUserId", h.getByUserId)
	// 修改
	authUserRouter.POST("/updateUser", h.updateUser)

}
