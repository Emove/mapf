package data

import "time"

const (
	InValidStatus = 0
	ValidStatus   = 1
)

const (
	DisableStatus = 0
	EnableStatus  = 1
)

const (
	IsDefaultFalse = 0
	IsDefaultTrue  = 1
)

const (
	IsDeletedFalse = 0
	IsDeletedTrue  = 1
)

type Model struct {
	ID         int       `gorm:"primaryKey"`
	CreateTime time.Time `gorm:"type:timestamp;autoCreateTime;not null"`
	UpdateTime time.Time `gorm:"type:timestamp;autoUpdateTime;not null"`
	IsDeleted  int       `gorm:"default:0;not null"`
}
