package common

import (
	"LiadminApi/modules"
	"LiadminApi/service"
	"LiadminApi/utils"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)



type HandlerCommon struct {

}

// 业务逻辑
var Rsp = &utils.Result{}

// login godoc
// @Summary 用户登录
// @Schemes
// @Description 用户登录接口，验证用户名和密码，返回访问令牌和刷新令牌
// @Tags user
// @Accept json
// @Produce json
// @Param request body modules.LoginRequest true "登录请求"
// @Success 200 {object} modules.LoginResponse
// @Router /login [post]
func (*HandlerCommon) login(ctx *gin.Context) {

	loginRequest := &modules.LoginRequest{}

	if err := ctx.ShouldBindJSON(loginRequest); err != nil {
		ctx.JSON(200, Rsp.Fail(400, "用户名或密码为空，请重新输入"))
		return
	}

	fmt.Println(loginRequest)

	user, err := service.GetByUserName(&modules.SysUserModule{UserName: loginRequest.UserName})

	if err != nil {
		ctx.JSON(200, Rsp.Fail(400, err.Error()))
		return
	}

	if !utils.VerifyPassword(loginRequest.PassWord, user.Salt, user.PassWord) {
		ctx.JSON(200, Rsp.Fail(400, "密码不正确请重新输入"))
		return
	}

	now := time.Now()
	user.LoginAt = &now
	// 更新用户信息到数据库
	if err := service.UpdateUser(user); err != nil {
		ctx.JSON(200, Rsp.Fail(400, "更新用户信息失败"))
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

// registerUser godoc
// @Summary 注册
// @Schemes
// @Description 注册
// @Tags user
// @Accept json
// @Produce json
// @Param request body modules.LoginRequest true "登录请求"
// @Success 200 string string 注册成功
// @Router /registerUser [post]
func (*HandlerCommon) registerUser(ctx *gin.Context) {

	registerRequest := &modules.SysUserModule{}

	if err := ctx.ShouldBindJSON(registerRequest); err != nil {
		ctx.JSON(200, Rsp.Fail(400, "注册信息不完整，请重新输入"))
		return
	}

	// 调用注册逻辑
	_, err := service.GetByUserName(&modules.SysUserModule{UserName: registerRequest.UserName})

	if err != nil {

		registerRequest.Salt = utils.GenerateSalt()
		// 加密
		registerRequest.PassWord = utils.HashPassword(registerRequest.PassWord, registerRequest.Salt)

		createError := service.CreateUser(registerRequest)

		if createError != nil {
			ctx.JSON(200, Rsp.Fail(400, err.Error()))
		} else {
			ctx.JSON(200, Rsp.Success(nil))
		}
	} else {
		ctx.JSON(200, Rsp.Fail(400, "用户名已存在"))
	}
}