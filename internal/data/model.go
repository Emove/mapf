package data

import "time"

const (
	ValidStatus   = 1
	InValidStatus = 2
)

const (
	EnableStatus  = 1
	DisableStatus = 2
)

const (
	IsDefaultTrue  = 1
	IsDefaultFalse = 2
)

const (
	IsDeletedTrue  = 1
	IsDeletedFalse = 2
)

type Model struct {
	ID         int       `gorm:"primaryKey"`
	CreateTime time.Time `gorm:"type:timestamp;autoCreateTime;not null"`
	UpdateTime time.Time `gorm:"type:timestamp;autoUpdateTime;not null"`
	IsDeleted  int       `gorm:"default:2;not null"`
	changes    map[string]interface{}
}

func (model *Model) Update(filedName string, value interface{}) {
	if model.changes == nil {
		model.changes = make(map[string]interface{})
	}
	model.changes[filedName] = value
}

func (model *Model) GetChanges() map[string]interface{} {
	if model.changes == nil {
		model.changes = make(map[string]interface{})
	}
	return model.changes
}

// IsValidValidStatusValue 判断是否是有效的有效状态码
func IsValidValidStatusValue(validStatus int) bool {
	if validStatus == ValidStatus || validStatus == InValidStatus {
		return true
	}
	return false
}

// IsValidEnableStatusValue 判断是否是有效的启用状态码
func IsValidEnableStatusValue(enableStatus int) bool {
	if enableStatus == EnableStatus || enableStatus == DisableStatus {
		return true
	}
	return false
}

// IsValidDefaultStatus 判断是否是有效的默认状态码
func IsValidDefaultStatus(isDefault int) bool {
	if isDefault == IsDefaultFalse || isDefault == IsDefaultTrue {
		return true
	}
	return false
}

func IsDeleted(isDelete int) bool {
	return isDelete == IsDeletedTrue
}

func IsDefault(isDefault int) bool {
	return isDefault == IsDefaultTrue
}

func IsEnable(enableStatus int) bool {
	return enableStatus == EnableStatus
}

func IsDisable(enableStatus int) bool {
	return DisableStatus == enableStatus
}
