package modules

import "gorm.io/gorm"

// 角色表
type SysRoleModule struct {
	gorm.Model
	RoleName string `gorm:"column:role_name;type:varchar(100);" json:"roleName"`
	Sort     string `gorm:"column:sort;type:int(4);" json:"sort"`           //显示顺序
	Status   string `gorm:"column:status;type:varchar(1);" json:"status"`   // '角色状态（0正常 1停用）',
	Remark   string `gorm:"column:remark;type:varchar(500);" json:"remark"` // '角色状态（0正常 1停用）',
}

func (table *SysRoleModule) TableName() string {
	return "sys_role"
}
