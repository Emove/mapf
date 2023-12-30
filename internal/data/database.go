package data

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
)

type Database struct {
	Host         string
	Port         string
	User         string
	Password     string
	DatabaseName string
}

func NewPostgresDB(config *Database, logger *log.Helper) (*gorm.DB, error) {
	port, err := strconv.ParseInt(config.Port, 10, 16)
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		config.Host, config.User, config.Password, config.DatabaseName, port)
	logger.Infow("database dsn", dsn)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁止自动创建外键
	})
}
