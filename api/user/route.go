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
	//注册
	r.POST("/register", h.registerUser)
	// 登陆
	r.POST("/login", h.login)
	// 用户list
	r.GET("/getByUserList", h.getByUserList)
	// 删除user
	r.DELETE("/deleteUserById", h.deleteUserById)
	// 根据id查询用户
	r.GET("/getByUserId", h.getByUserId)
	// 修改
	r.POST("/updateUser", h.updateUser)
}
