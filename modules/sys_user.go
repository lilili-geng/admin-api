package modules

import (
	"time"

	"gorm.io/gorm"
)

// 用户信息表
type SysUserModule struct {
	ID         uint           `gorm:"primarykey;" json:"id"`
	UserName   string         `gorm:"column:username;type:varchar(50);" json:"userName"`
	PassWord   string         `gorm:"column:password;type:varchar(36);" json:"passWord"`
	Phone      string         `gorm:"column:phone;type:varchar(20);" json:"phone"`
	Avatar     string         `gorm:"column:avatar;type:varchar(255);" json:"avatar"`
	Email      string         `gorm:"column:email;type:varchar(30);" valid:"email" json:"email"`
	Salt       string         `gorm:"column:salt;" json:"salt"`
	LoginAt    *time.Time     `gorm:"column:login_at;" json:"loginAt"`
	LoginOutAt *time.Time     `gorm:"column:login_out_at;" json:"loginOutAt"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	IsLogOut   int           `gorm:"column:is_log_out;" json:"isLogOut"`
}

// login request
type LoginRequest struct {
	UserName string `json:"userName"`
	PassWord string `json:"password"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

// userList
type UserPagination struct {
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

type UserPaginationResponse struct {
	Page     int             `json:"page"`
	PageSize int             `json:"pageSize"`
	Total    int             `json:"total"`
	List     []SysUserModule `json:"list"`
}

func (table *SysUserModule) TableName() string {
	return "sys_user"
}
