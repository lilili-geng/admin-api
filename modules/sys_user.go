package modules

import (
	"time"

	"gorm.io/gorm"
)

// 用户信息表
type SysUserModule struct {
	gorm.Model
	UserName     string `gorm:"column:username;type:varchar(50);" json:"userName"`
	PassWord     string `gorm:"column:password;type:varchar(36);" json:"passWord"`
	Phone        string `gorm:"column:phone;type:varchar(20);" json:"phone"`
	Avatar       string `gorm:"column:avatar;type:varchar(255);" json:"avatar"`
	Email        string `gorm:"column:avatar;type:varchar(30);" valid:"email" json:"email"`
	Salt         string
	LoginTime    *time.Time
	LoginOutTime *time.Time `gorm:"column:login_out_time" json:"loginOutTime"`
	IsLogOut     bool
}

// login request
type LoginRequest struct {
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

func (table *SysUserModule) TableName() string {
	return "sys_user"
}
