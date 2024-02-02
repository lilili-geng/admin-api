package user

import (
	"LiadminApi/modules"
	"LiadminApi/service"
	"LiadminApi/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
}

// 业务逻辑
var Rsp = &utils.Result{}

// getByUserList
func (*HandlerUser) getByUserList(ctx *gin.Context) {
	user := &modules.UserPagination{}
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	user.UserName = ctx.Query("userName")
	user.Email = ctx.Query("email")
	user.Page = page
	user.PageSize = pageSize

	result, err := service.GetUserList(user)

	if err != nil {
		ctx.JSON(200, Rsp.Fail(400, err.Error()))
		return
	}

	ctx.JSON(200, Rsp.Success(result))
}

// deleteUserById
func (*HandlerUser) deleteUserById(ctx *gin.Context) {

	userIds := strings.Split(ctx.Query("userId"), ",")

	if len(userIds) == 0 {
		ctx.JSON(400, Rsp.Fail(400, "UserIds 数组为空"))
		return
	}

	var ids []int

	for _, id := range userIds {
		userId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(400, Rsp.Fail(400, "无效的用户ID"))
			return
		}
		ids = append(ids, userId)
	}

	fmt.Println(userIds)

	err := service.DeleteUserById(ids)

	if err != nil {
		ctx.JSON(500, Rsp.Fail(500, err.Error()))
		return
	}

	ctx.JSON(200, Rsp.Success(nil))
}

// getByUserId
func (*HandlerUser) getByUserId(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Query("id"))

	if err != nil {
		// 处理错误，例如返回错误响应
		ctx.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	result, err := service.GetByUserId(int64(id))

	if err != nil {
		ctx.JSON(400, Rsp.Fail(400, err.Error()))
		return
	}

	ctx.JSON(200, Rsp.Success(result))
}

// updateUser
func (*HandlerUser) updateUser(ctx *gin.Context) {

	user := &modules.SysUserModule{}

	if err := ctx.ShouldBindJSON(user); err != nil {
		ctx.JSON(200, Rsp.Fail(400, err.Error()))
		return
	}

	err := service.UpdateUser(user)
	if err != nil {
		ctx.JSON(200, Rsp.Fail(400, err.Error()))
		return
	}

	ctx.JSON(200, Rsp.Success(nil))
}
