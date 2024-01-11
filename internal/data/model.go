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
