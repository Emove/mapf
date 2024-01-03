package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
	"mapf/app/warehouse/internal/conf"
	"os"
	"testing"
)

const configPath = "../../configs"

func newConfig() *conf.Data {
	c := config.New(
		config.WithSource(
			file.NewSource(configPath),
		),
	)
	defer func(c config.Config) {
		_ = c.Close()
	}(c)

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	return bc.Data
}

func newLogger() log.Logger {
	return log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.name", "warehouse",
	)
}

func TestNewData(t *testing.T) {
	type args struct {
		c      *conf.Data
		logger log.Logger
	}
	tests := []struct {
		name    string
		args    args
		want    func(got *Data)
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "normal",
			args: args{
				c:      newConfig(),
				logger: newLogger(),
			},
			want: func(got *Data) {
				assert.NotNilf(t, got.DB(context.Background()), "DB return nil")
				assert.NotNil(t, got.db, "db is nil")
				assert.NotNil(t, got.rds, "rds is nil")
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf(i[0].(string), i[1:]...)
					return false
				}
				return true
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := NewData(tt.args.c, tt.args.logger)
			if !tt.wantErr(t, err, fmt.Sprintf("NewData(%v, %v), err: %v", tt.args.c, tt.args.logger, err)) {
				return
			}
			tt.want(got)
		})
	}
}
