package data

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"mapf/app/warehouse/internal/biz"
	"mapf/internal/data"
	"testing"
)

func newWarehouseRepo(t *testing.T) biz.WarehouseRepo {
	logger := newLogger()
	data, _, err := NewData(newConfig(), logger)
	assert.Nil(t, err, "new data error")
	repo, err := NewWarehouseRepo(data, logger)
	assert.Nil(t, err, "new warehouse repo error")
	return repo
}

func Test_warehouseRepo_CreateWarehouse(t *testing.T) {
	type args struct {
		ctx       context.Context
		warehouse *biz.Warehouse
	}
	tests := []struct {
		name             string
		args             args
		want             func(*biz.Warehouse)
		errFn            func(err error)
		deletedAfterTest bool
	}{
		{
			name: "first",
			args: args{ctx: context.Background(), warehouse: &biz.Warehouse{
				Name: "test",
			}},
			want: func(warehouse *biz.Warehouse) {
				assert.NotNilf(t, warehouse.ID, "warehouse ID is nil")
				assert.Equal(t, warehouse.Name, "test", "warehouse name not equal")
				assert.Equal(t, warehouse.IsDefault, data.IsDefaultFalse, "is a default warehouse")
				assert.Equal(t, warehouse.Status, data.EnableStatus, "warehouse status not equal to Enable")
				assert.Equal(t, warehouse.IsDeleted, data.IsDeletedFalse, "warehouse is deleted")
			},
			errFn: func(err error) {
				if err != nil {
					assert.ErrorIs(t, err, gorm.ErrDuplicatedKey, "got an unexpected error")
				}
			},
			deletedAfterTest: true,
		},
	}
	repo := newWarehouseRepo(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.CreateWarehouse(tt.args.ctx, tt.args.warehouse)
			tt.errFn(err)
			tt.want(got)
		})
	}
}

func Test_warehouseRepo_GetWarehouseByName(t *testing.T) {
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    func(t assert.TestingT, got *biz.Warehouse)
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "normal",
			args: args{
				ctx:  context.Background(),
				name: "test",
			},
			want: func(t assert.TestingT, got *biz.Warehouse) {
				assert.NotNilf(t, got, "warehouse is nil")
				assert.Equal(t, got.Name, "test", "warehouse name not equal to test")
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf(i[0].(string), i[1:]...)
					return false
				}
				return true
			},
		},
		{
			name: "not found",
			args: args{
				ctx:  context.Background(),
				name: "213123214",
			},
			want: func(t assert.TestingT, got *biz.Warehouse) {
				assert.Nil(t, got, "warehouse is not nil")
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
	repo := newWarehouseRepo(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.GetWarehouseByName(tt.args.ctx, tt.args.name)
			if !tt.wantErr(t, err, fmt.Sprintf("GetWarehouseByName(%v, %v), err: %v", tt.args.ctx, tt.args.name, err)) {
				return
			}
			tt.want(t, got)
		})
	}
}

func Test_warehouseRepo_GetWarehouseById(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		args    args
		want    func(t assert.TestingT, got *biz.Warehouse)
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "normal",
			args: args{
				ctx: context.Background(),
				id:  9,
			},
			want: func(t assert.TestingT, got *biz.Warehouse) {
				assert.NotNilf(t, got, "warehouse is nil")
				assert.Equal(t, got.ID, 9, "warehouse name not equal to want")
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					t.Errorf(i[0].(string), i[1:]...)
					return false
				}
				return true
			},
		},
		{
			name: "not found",
			args: args{
				ctx: context.Background(),
				id:  12312,
			},
			want: func(t assert.TestingT, got *biz.Warehouse) {
				assert.Nil(t, got, "warehouse is not nil")
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
	repo := newWarehouseRepo(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.GetWarehouseById(tt.args.ctx, tt.args.id)
			if !tt.wantErr(t, err, fmt.Sprintf("GetWarehouseById(%v, %v), err: %v", tt.args.ctx, tt.args.id, err)) {
				return
			}
			tt.want(t, got)
		})
	}
}
