package service

import (
	"LiadminApi/modules"
	"LiadminApi/utils/db"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
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

	list := &modules.UserPaginationResponse{
		Page:     user.Page,
		PageSize: user.PageSize,
		Total:    int(total),
		List:     userList,
	}
	return list, nil
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

// 查询用户个人信息
// getUserInfo
func GetUserInfo(name string) (*modules.GetUserInfoResponse, error) {
	userRole := &modules.SysUserRole{}
	userInfo := &modules.GetUserInfoResponse{}

	user, err := GetByUserName(&modules.SysUserModule{UserName: name})

	if err != nil {
		return nil, err
	}

	// 查询用户角色关联表
	if err := db.DB.Preload("User").Preload("Role").Where("user_id = ?", user.ID).First(&userRole).Error; err != nil {
		// 处理查询错误
		return nil, err
	}

	userInfo.SysUserModule = userRole.User

	userInfo.RoleName = userRole.Role.RoleName
	userInfo.Remark = userRole.Role.Remark
	userInfo.Sort = userRole.Role.Sort
	userInfo.RoleStatus = userRole.Role.Status

	fmt.Println("GetUserInfo", userInfo)

	return userInfo, nil
}
