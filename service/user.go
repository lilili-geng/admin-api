package service

import (
	"LiadminApi/modules"
	"LiadminApi/utils/db"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// 查询用户名和密码
func GetByUserNamePassword(userName string, passWord string) (*modules.SysUserModule, error) {
	data := &modules.SysUserModule{}
	
	// 查询的时候排出软删后的
	err := db.DB.Where("username = ? AND password = ?", userName, passWord).Where("deleted_at IS NULL").First(&data).Error
	
	if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, fmt.Errorf("用户名或密码错误，请重新输入")
			}
			return nil, err
	}
	
	return data, nil
}

