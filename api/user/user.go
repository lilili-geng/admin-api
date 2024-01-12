package user

import "github.com/gin-gonic/gin"

type HandlerUser struct {
}

// 业务逻辑
func (*HandlerUser) getCaptcha(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "login",
	})
}
