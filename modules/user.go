package modules

import (
	"time"

	"gorm.io/gorm"
)

type UserModule struct {
	gorm.Model
	Identity      string
	Name          string
	PassWord      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	ClientIp      string
	ClientPort    string
	Salt          string
	LoginTime     *time.Time
	HeartbeatTime *time.Time
	LoginOutTime  *time.Time `gorm:"column:login_out_time" json:"login_out_time"`
	IsLogOut      bool
	DeviceInfo    string
}

func (table *UserModule) TableName() string {
	return "user"
}
