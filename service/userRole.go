package service

import (
	"LiadminApi/modules"
	"LiadminApi/utils/db"
)

// 插入
func CreateUserRole(userId int, roleId int) error {
	return db.DB.Create(&modules.SysUserRole{UserID: uint(userId), RoleID: uint(roleId)}).Error
}

func GetByRoleUserId(userId int64) (*modules.SysUserRole, error) {
	userRole := &modules.SysUserRole{}
	if err := db.DB.Where("user_id = ?", userId).First(userRole).Error; err != nil {
		return nil, err
	}
	return userRole, nil
}

func UpdateUserRole(userId int64, roleId int64) error {
	return db.DB.Model(&modules.SysUserRole{}).Where("user_id = ?", userId).Update("role_id", roleId).Error
}
