package service

import (
	"LiadminApi/modules"
	"LiadminApi/utils/db"
	"errors"

	"github.com/jinzhu/gorm"
)

// GetByRoleId
func GetByRoleId(id int64) (*modules.SysRoleModule, error) {
	data := &modules.SysRoleModule{}

	err := db.DB.Where("id = ?", id).First(data).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	return data, nil
}

func GetRoleList(role *modules.RolePagination) (*modules.RolePaginationResponse, error) {
	var roleList []modules.SysRoleModule
	var result *gorm.DB
	var total int64

	baseQuery := db.DB

	if role.RoleName != "" {
		baseQuery = baseQuery.Where("role_name LIKE ?", "%"+role.RoleName+"%")
	}

	if role.Status != "" {
		baseQuery = baseQuery.Where("status LIKE ?", "%"+role.Status+"%")
	}

	result = baseQuery.Limit(role.PageSize).Offset((role.Page - 1) * role.PageSize).Find(&roleList)

	baseQuery.Model(&modules.SysRoleModule{}).Count(&total)

	if result == nil {
		return nil, errors.New("result is nil")
	}

	list := &modules.RolePaginationResponse{
		Page:     role.Page,
		PageSize: role.PageSize,
		Total:    int(total),
		List:     roleList,
	}

	return list, nil
}

// 修改
func UpdateRole(role *modules.SysRoleModule) error {

	result := db.DB.Model(&role).Update(modules.SysRoleModule{
		RoleName: role.RoleName,
		Status:   role.Status,
		Remark:   role.Remark,
		Sort:     role.Sort,
	}).Error

	return result
} 

