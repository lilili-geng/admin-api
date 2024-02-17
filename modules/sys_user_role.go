package modules

import (
	_ "gorm.io/gorm"
)


// 用户和角色关联表
type SysUserRole struct {
	UserID uint `gorm:"column:user_id;primaryKey" json:"userID"` // 用户ID
	RoleID uint `gorm:"column:role_id;primaryKey" json:"roleID"` // 角色ID
}

func (SysUserRole) TableName() string {
	return "sys_user_role"
}
