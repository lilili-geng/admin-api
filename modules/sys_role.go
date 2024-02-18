package modules

import (
	"time"

	"gorm.io/gorm"
)

// 角色表
type SysRoleModule struct {
	ID        uint           `gorm:"primarykey;" json:"id"`
	RoleName  string         `gorm:"column:role_name;type:varchar(100);" json:"roleName"` // 权限名字
	Sort      string         `gorm:"column:sort;type:int(4);" json:"sort"`                // 显示顺序
	Status    string         `gorm:"column:status;type:varchar(1);" json:"roleStatus"`    // 角色状态（0正常 1停用）,
	Remark    string         `gorm:"column:remark;type:varchar(500);" json:"remark"`      // 备注,
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}

// role
type RolePagination struct {
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	RoleName string `json:"roleName"`
	Status   string `json:"roleStatus"`
}

type RolePaginationResponse struct {
	Page     int             `json:"page"`
	PageSize int             `json:"pageSize"`
	Total    int             `json:"total"`
	List     []SysRoleModule `json:"list"`
}

func (table *SysRoleModule) TableName() string {
	return "sys_role"
}
