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

// getUserList
func GetUserList(user *modules.UserPagination) (*modules.UserPaginationResponse, error) {

	fmt.Println(user, "user11111")

	var userList []modules.SysUserModule

	var result *gorm.DB

	if user.Email != "" && user.UserName != "" {
		result = db.DB.Where("username LIKE ? OR email LIKE ?", "%"+user.UserName+"%", "%"+user.Email+"%").Limit(user.PageSize).Offset((user.Page - 1) * user.PageSize).Find(&userList)
	} else {
		result = db.DB.Limit(user.PageSize).Offset((user.Page - 1) * user.PageSize).Find(&userList)
	}

	if user.UserName != "" {
		result = db.DB.Where("username LIKE ? ", "%"+user.UserName+"%").Limit(user.PageSize).Offset((user.Page - 1) * user.PageSize).Find(&userList)
	}

	if user.Email != "" {
		result = db.DB.Where("email LIKE ?", "%"+user.Email+"%").Limit(user.PageSize).Offset((user.Page - 1) * user.PageSize).Find(&userList)
	}

	if result == nil {
		return nil, errors.New("result is nil")
	}

	pagination := &modules.UserPaginationResponse{
		Page:     user.Page,
		PageSize: user.PageSize,
		Total:    len(userList),
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
