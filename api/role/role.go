package role

import (
	"LiadminApi/modules"
	"LiadminApi/service"
	"LiadminApi/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HandlerRole struct {
}

// 业务逻辑
var Rsp = &utils.Result{}

// getrolelist
func (*HandlerRole) getRoleList(ctx *gin.Context) {
	role := &modules.RolePagination{}
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	role.Status = ctx.Query("roleStatus")
	role.RoleName = ctx.Query("roleName")
	role.Page = page
	role.PageSize = pageSize

	result, err := service.GetRoleList(role)

	if err != nil {
		ctx.JSON(200, Rsp.Fail(400, err.Error()))
		return
	}

	ctx.JSON(200, Rsp.Success(result))
}

// updateRole
func (*HandlerRole) updateRole(ctx *gin.Context) {
	role := &modules.SysRoleModule{}

	if err := ctx.ShouldBindJSON(role); err != nil {
		ctx.JSON(200, Rsp.Fail(400, err.Error()))
		return
	}

	if role.ID == 0 {
		ctx.JSON(200, Rsp.Fail(401, "角色id不存在"))
		return
	}

	err := service.UpdateRole(role)

	if err != nil {
		ctx.JSON(200, Rsp.Fail(401, err.Error()))
		return
	}

	ctx.JSON(200, Rsp.Success("修改成功"))
}
