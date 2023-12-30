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
	ID          uint      `gorm:"primaryKey"`
	CreatedTime time.Time `gorm:"type:timestamp;autoCreateTime"`
	UpdatedTime time.Time `gorm:"type:timestamp;autoUpdateTime"`
	IsDeleted   int       `gorm:"default:0"`
}
