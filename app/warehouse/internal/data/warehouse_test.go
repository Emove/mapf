package data

import (
	"context"
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
