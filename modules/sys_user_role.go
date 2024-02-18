package modules

import (
	_ "gorm.io/gorm"
)

// 用户和角色关联表
type SysUserRole struct {
	UserID uint          `gorm:"column:user_id;primaryKey" json:"userID"` // 用户ID
	RoleID uint          `gorm:"column:role_id;primaryKey" json:"roleID"` // 角色ID
	User   SysUserModule `gorm:"foreignKey:UserID"`                       // 定义用户外键
	Role   SysRoleModule `gorm:"foreignKey:RoleID"`                       // 定义角色外键
}

// 根据jwt keyname 查询 当前用户个人信息
type GetUserInfoResponse struct {
	SysUserModule
	RoleName   string `json:"roleName"`
	Sort       string `json:"sort"`
	RoleStatus string `json:"roleStatus"`
	Remark     string `json:"remark"`
}
	
type CreateUserRoleRequest struct {
	SysUserModule
	RoleID uint `gorm:"column:role_id;primaryKey" json:"roleID"`
}

func (SysUserRole) TableName() string {
	return "sys_user_role"
}
