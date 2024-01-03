package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"strconv"
	"time"
)

type Database struct {
	Host         string
	Port         string
	User         string
	Password     string
	DatabaseName string
}

func NewMysqlDB(config *Database, logger log.Logger, helper *log.Helper) (*gorm.DB, error) {
	port, err := strconv.ParseInt(config.Port, 10, 16)
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, port, config.DatabaseName)
	helper.Infow("database_dsn", dsn)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   NewGormLogger(logger),
		TranslateError:                           true, // 将数据库错误转换成gorm中定义的错误
		DisableForeignKeyConstraintWhenMigrating: true, // 禁止自动创建外键
		SkipDefaultTransaction:                   true, // 默认不开启事务
	})
}

func NewPostgresDB(config *Database, logger log.Logger, helper *log.Helper) (*gorm.DB, error) {
	port, err := strconv.ParseInt(config.Port, 10, 16)
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		config.Host, config.User, config.Password, config.DatabaseName, port)
	helper.Infow("database_dsn", dsn)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                                   NewGormLogger(logger),
		TranslateError:                           true, // 将数据库错误转换成gorm中定义的错误
		DisableForeignKeyConstraintWhenMigrating: true, // 禁止自动创建外键
		SkipDefaultTransaction:                   true, // 默认不开启事务
	})
}

type GormLogger struct {
	logger        log.Logger
	level         logger.LogLevel
	helper        *log.Helper
	slowThreshold time.Duration
}

func NewGormLogger(logger log.Logger) *GormLogger {
	return &GormLogger{logger: logger, helper: log.NewHelper(logger), slowThreshold: 200 * time.Millisecond}
}

func (gl *GormLogger) LogMode(logLevel logger.LogLevel) logger.Interface {
	if logLevel == logger.Silent {
		return logger.Default.LogMode(logger.Silent)
	}
	return gl
}
func (gl *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if gl.level == logger.Info {
		gl.helper.WithContext(ctx).Infow("line", utils.FileWithLineNum(), fmt.Sprintf(msg, data...))
	}
}
func (gl *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if gl.level >= logger.Warn {
		gl.helper.WithContext(ctx).Warnw("line", utils.FileWithLineNum(), fmt.Sprintf(msg, data...))
	}
}
func (gl *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if gl.level >= logger.Error {
		gl.helper.WithContext(ctx).Errorw("line", utils.FileWithLineNum(), fmt.Sprintf(msg, data...))
	}
}
func (gl *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if gl.level <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && gl.level >= logger.Error && !errors.Is(err, gorm.ErrRecordNotFound):
		sql, rows := fc()
		if rows == -1 {
			gl.helper.Errorw(
				"line", utils.FileWithLineNum(),
				"elapsed", fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6),
				"sql", sql, "error", err)
		} else {
			gl.helper.Errorw(
				"line", utils.FileWithLineNum(),
				"elapsed", fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6),
				"sql", sql, "rows", rows, "error", err)
		}
	case elapsed > gl.slowThreshold && gl.slowThreshold != 0 && gl.level >= logger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", gl.slowThreshold)
		if rows == -1 {
			gl.helper.Warnw(
				"line", utils.FileWithLineNum(),
				"slow", slowLog, "elapsed",
				fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6),
				"sql", sql)
		} else {
			gl.helper.Warnw(
				"line", utils.FileWithLineNum(),
				"slow", slowLog, "elapsed",
				fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6),
				"rows", rows,
				"sql", sql)
		}
	case gl.level == logger.Info:
		sql, rows := fc()
		if rows == -1 {
			gl.helper.Infow(
				"line", utils.FileWithLineNum(),
				"elapsed", fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6),
				"sql", sql)
		} else {
			gl.helper.Infow(
				"line", utils.FileWithLineNum(),
				"elapsed", fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6),
				"rows", rows,
				"sql", sql)
		}
	}
}
