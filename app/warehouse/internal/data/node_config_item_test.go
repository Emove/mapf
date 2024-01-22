package data

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"mapf/app/warehouse/internal/biz"
	"mapf/internal/data"
	"testing"
)

func newNodeConfigItemRepo(t *testing.T) biz.NodeConfigItemRepo {
	logger := newLogger()
	config := newConfig()
	//config.Database.LogLevel = string(data.LogLevelSilent)
	data, _, err := NewData(config, logger)
	assert.Nil(t, err, "new data error")
	repo, err := NewNodeConfigItemRepo(data, logger)
	assert.Nil(t, err, "new Node repo error")
	return repo
}

func Test_nodeConfigItemRepo_CreateNodeConfigItem(t *testing.T) {
	type args struct {
		ctx  context.Context
		item *biz.NodeConfigItem
	}
	boolOptionalValues, _ := json.Marshal([]string{"true", "false"})
	tests := []struct {
		name    string
		args    args
		want    func(*biz.NodeConfigItem)
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "normal",
			args: args{
				ctx:  context.Background(),
				item: &biz.NodeConfigItem{WarehouseId: 9, NodeTypeId: 18, Code: "rotatable", Name: "是否可旋转", ValueType: "bool", DefaultValue: "true", OptionalValues: string(boolOptionalValues)},
			},
			want: func(item *biz.NodeConfigItem) {
				assert.NotNil(t, item.ID)
			},
			wantErr: ignoreDuplicatedKeyError,
		},
	}
	repo := newNodeConfigItemRepo(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.CreateNodeConfigItem(tt.args.ctx, tt.args.item)
			if !tt.wantErr(t, err, fmt.Sprintf("CreateNodeConfigItem(%v, %v)", tt.args.ctx, tt.args.item)) {
				return
			}
			tt.want(got)
		})
	}
}

func Test_nodeConfigItemRepo_UpdateNodeConfigItem(t *testing.T) {
	type args struct {
		ctx  context.Context
		item *biz.NodeConfigItem
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "update_code",
			args: args{ctx: context.Background(), item: setCode(1, "rotate")},
		},
		{
			name: "update_name",
			args: args{ctx: context.Background(), item: setName(1, "可旋转")},
		},
		{
			name: "update_value_type",
			args: args{ctx: context.Background(), item: setValueType(1, "boolean")},
		},
		{
			name: "update_default_value",
			args: args{ctx: context.Background(), item: setDefaultValue(1, "False")},
		},
		{
			name: "update_optional_values",
			args: args{ctx: context.Background(), item: setOptionalValues(1, "True", "False")},
		},
		{
			name: "rollback",
			args: args{ctx: context.Background(), item: setAll(1, "rotatable", "是否可旋转", "bool", "true", "true", "false")},
		},
	}
	repo := newNodeConfigItemRepo(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := repo.UpdateNodeConfigItem(tt.args.ctx, tt.args.item)
			if !defaultWantErrFunc(t, err, fmt.Sprintf("UpdateNodeConfigItem(%v, %v)", tt.args.ctx, tt.args.item)) {
				return
			}
		})
	}
}

func setCode(id int, code string) *biz.NodeConfigItem {
	item := &biz.NodeConfigItem{Model: &data.Model{ID: id}}
	item.SetCode(code)
	return item
}

func setName(id int, name string) *biz.NodeConfigItem {
	item := &biz.NodeConfigItem{Model: &data.Model{ID: id}}
	item.SetName(name)
	return item
}

func setValueType(id int, valueType string) *biz.NodeConfigItem {
	item := &biz.NodeConfigItem{Model: &data.Model{ID: id}}
	_ = item.SetValueType(valueType)
	return item
}

func setDefaultValue(id int, defaultValue string) *biz.NodeConfigItem {
	item := &biz.NodeConfigItem{Model: &data.Model{ID: id}}
	_ = item.SetDefaultValue(defaultValue)
	return item
}

func setOptionalValues(id int, optionalValues ...string) *biz.NodeConfigItem {
	item := &biz.NodeConfigItem{Model: &data.Model{ID: id}}
	ovs, _ := json.Marshal(optionalValues)
	_ = item.SetOptionalValues(string(ovs))
	return item
}

func setAll(id int, code, name, valueType, defaultValue string, optionalValues ...string) *biz.NodeConfigItem {
	item := &biz.NodeConfigItem{Model: &data.Model{ID: id}}
	item.SetCode(code)
	item.SetName(name)
	_ = item.SetValueType(valueType)
	_ = item.SetDefaultValue(defaultValue)
	ovs, _ := json.Marshal(optionalValues)
	_ = item.SetOptionalValues(string(ovs))
	return item
}

func Test_nodeConfigItemRepo_GetNodeConfigItemByWarehouseIdAndNodeTypeId(t *testing.T) {
	type args struct {
		ctx         context.Context
		warehouseId int
		nodeTypeId  *int
	}
	var existedNodeTypeId = 18
	var notExistedNodeTypeId = 17
	var tests = []struct {
		name      string
		args      args
		wantItems func([]*biz.NodeConfigItem)
		wantErr   assert.ErrorAssertionFunc
	}{
		{
			name: "select_by_warehouse_id",
			args: args{ctx: context.Background(), warehouseId: 9},
			wantItems: func(items []*biz.NodeConfigItem) {
				assert.Greater(t, len(items), 0)
				for _, item := range items {
					assert.Equal(t, 9, item.WarehouseId)
				}
			},
			wantErr: defaultWantErrFunc,
		},
		{
			name: "select_by_warehouse_id_and_node_type_id_normal",
			args: args{ctx: context.Background(), warehouseId: 9, nodeTypeId: &existedNodeTypeId},
			wantItems: func(items []*biz.NodeConfigItem) {
				assert.Greater(t, len(items), 0)
				for _, item := range items {
					assert.Equal(t, 9, item.WarehouseId)
					assert.Equal(t, existedNodeTypeId, item.NodeTypeId)
				}
			},
			wantErr: defaultWantErrFunc,
		},
		{
			name: "select_by_warehouse_id_and_node_type_id_not_found",
			args: args{ctx: context.Background(), warehouseId: 9, nodeTypeId: &notExistedNodeTypeId},
			wantItems: func(items []*biz.NodeConfigItem) {
				assert.Equal(t, len(items), 0)
			},
			wantErr: defaultWantErrFunc,
		},
	}
	repo := newNodeConfigItemRepo(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItems, err := repo.GetNodeConfigItemByWarehouseIdAndNodeTypeId(tt.args.ctx, tt.args.warehouseId, tt.args.nodeTypeId)
			if !tt.wantErr(t, err, fmt.Sprintf("GetNodeConfigItemByWarehouseIdAndNodeTypeId(%v, %v, %v)", tt.args.ctx, tt.args.warehouseId, tt.args.nodeTypeId)) {
				return
			}
			tt.wantItems(gotItems)
		})
	}
}
