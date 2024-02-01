package service

import (
	"LiadminApi/modules"
	"LiadminApi/utils/db"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// 查询用户名
func GetByUserName(user *modules.SysUserModule) (*modules.SysUserModule, error) {
	data := &modules.SysUserModule{}

	// 查询的时候排除软删后的
	err := db.DB.Where("username = ? AND deleted_at IS NULL", user.UserName).First(&data).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("该用户未注册")
		}
		return nil, err
	}

	return data, nil
}

// CreateUser 创建用户
func CreateUser(user *modules.SysUserModule) error {
	err := db.DB.Create(&user).Error
	return err
}

func GetUserList(user *modules.UserPagination) (*modules.UserPaginationResponse, error) {
	fmt.Println(user, "user11111")

	var userList []modules.SysUserModule
	var result *gorm.DB
	var total int64

	baseQuery := db.DB

	if user.UserName != "" {
		baseQuery = baseQuery.Where("username LIKE ?", "%"+user.UserName+"%")
	}

	if user.Email != "" {
		baseQuery = baseQuery.Where("email LIKE ?", "%"+user.Email+"%")
	}

	result = baseQuery.Limit(user.PageSize).Offset((user.Page - 1) * user.PageSize).Find(&userList)

	baseQuery.Model(&modules.SysUserModule{}).Count(&total)

	if result == nil {
		return nil, errors.New("result is nil")
	}

	pagination := &modules.UserPaginationResponse{
		Page:     user.Page,
		PageSize: user.PageSize,
		Total:    int(total),
		List:     userList,
	}
	return pagination, nil
}

// UpdateUser
func UpdateUser(user *modules.SysUserModule) error {
	result := db.DB.Model(&user).Updates(modules.SysUserModule{
		UserName:   user.UserName,
		PassWord:   user.PassWord,
		Phone:      user.Phone,
		Avatar:     user.Avatar,
		Email:      user.Email,
		LoginAt:    user.LoginAt,
		LoginOutAt: user.LoginOutAt,
		IsLogOut:   user.IsLogOut,
	}).Error

	return result
}

// DeleteUserById
func DeleteUserById(userIds []int) error {
	fmt.Println(userIds, "userIds")
	for _, userID := range userIds {
		if result := db.DB.Delete(&modules.SysUserModule{}, userID); result.Error != nil {
			return result.Error
		}
	}
	return nil
}

// GetByUserId
func GetByUserId(userId int64) (*modules.SysUserModule, error) {
	data := &modules.SysUserModule{}

	err := db.DB.Where("id = ?", userId).First(data).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	return data, nil
}
