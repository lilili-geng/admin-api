package user

import (
	"LiadminApi/modules"
	"LiadminApi/service"
	"LiadminApi/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
}

// 业务逻辑
var Rsp = &utils.Result{}

// @BasePath /login
// @Summary 登陆
// Login
// @Tags user
// @param name body modules.UserLoginRequest false "用户名"
// @param passWord body modules.UserLoginRequest false "密码"
// @Success 200 {string}  json{"code","data"}
// @Router /login [post]
func (*HandlerUser) login(ctx *gin.Context) {

	loginRequest := &modules.LoginRequest{}

	if err := ctx.ShouldBindJSON(loginRequest); err != nil {
		ctx.JSON(200, Rsp.Fail(400, "用户名或密码为空，请重新输入"))
		return
	}

	user, err := service.GetByUserNamePassword(loginRequest.UserName, loginRequest.PassWord)

	if err != nil {
		ctx.JSON(200, Rsp.Fail(400, err.Error()))
		return
	}

	token, err := utils.GenerateToken(user.ID, user.UserName, utils.TokenExpire)

	if err != nil {
		ctx.JSON(200, Rsp.Fail(400, err.Error()))
		fmt.Println("token")
		return
	}

	refreshToken, err := utils.GenerateToken(user.ID, user.UserName, utils.RefreshTokenExpire)

	if err != nil {
		ctx.JSON(200, Rsp.Fail(400, err.Error()))
		fmt.Println("refreshToken")
		return
	}

	ctx.JSON(200, Rsp.Success(&modules.LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
	}))

	fmt.Println(user)

}
