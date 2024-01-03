package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"mapf/app/warehouse/internal/biz"
	"mapf/internal/data"
	"testing"
)

func newNodeTypeRepo(t *testing.T) biz.NodeTypeRepo {
	logger := newLogger()
	data, _, err := NewData(newConfig(), logger)
	assert.Nil(t, err, "new data error")
	repo, err := NewNodeTypeRepo(data, logger)
	assert.Nil(t, err, "new NodeType repo error")
	return repo
}

func Test_nodeTypeRepo_CreateNodeType(t *testing.T) {
	type args struct {
		ctx      context.Context
		nodeType *biz.NodeType
	}
	tests := []struct {
		name    string
		args    args
		want    func(want, got *biz.NodeType)
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "first",
			args: args{
				ctx: context.Background(),
				nodeType: &biz.NodeType{
					WarehouseId: 1,
					Code:        "RAIL",
					Name:        "巷道节点",
				},
			},
			want: func(want, got *biz.NodeType) {
				assert.EqualValuesf(t, want, got, "want: %v, got: %v", want, got)
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					if errors.Is(err, gorm.ErrDuplicatedKey) {
						return true
					}
					t.Errorf(i[0].(string), i[1:]...)
					return false
				}
				return true
			},
		},
		{
			name: "default_node_type",
			args: args{
				ctx: context.Background(),
				nodeType: &biz.NodeType{
					WarehouseId: 1,
					Code:        "CHARGE",
					Name:        "充电站",
					IsDefault:   data.IsDefaultTrue,
				},
			},
			want: func(want, got *biz.NodeType) {
				assert.EqualValuesf(t, want, got, "want: %v, got: %v", want, got)
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					if errors.Is(err, gorm.ErrDuplicatedKey) {
						return true
					}
					t.Errorf(i[0].(string), i[1:]...)
					return false
				}
				return true
			},
		},
	}
	repo := newNodeTypeRepo(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.CreateNodeType(tt.args.ctx, tt.args.nodeType)
			if !tt.wantErr(t, err, fmt.Sprintf("CreateNodeType(%v, %v), error: %v", tt.args.ctx, tt.args.nodeType, err)) {
				return
			}
			tt.want(tt.args.nodeType, got)
		})
	}
}

func Test_nodeTypeRepo_BatchCreateNodeTypes(t *testing.T) {
	type args struct {
		ctx       context.Context
		nodeTypes []*biz.NodeType
	}
	tests := []struct {
		name    string
		args    args
		want    func(t assert.TestingT, want, got []*biz.NodeType)
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "normal",
			args: args{
				ctx: context.Background(),
				nodeTypes: []*biz.NodeType{
					{
						WarehouseId: 1,
						Code:        "PASSAGE",
						Name:        "通道节点",
					},
					{
						WarehouseId: 1,
						Code:        "WORKSTATION",
						Name:        "工作站",
					},
				},
			},
			want: func(t assert.TestingT, want, got []*biz.NodeType) {
				if len(want) != len(got) {
					t.Errorf("len of want: %d, len of got: %d", len(want), len(got))
				}
				for i := 0; i < len(want); i++ {
					assert.EqualValuesf(t, want[i], got[i], "want: %v, got: %v", want[i], got[i])
				}
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if err != nil {
					if errors.Is(err, gorm.ErrDuplicatedKey) {
						return true
					}
					t.Errorf(i[0].(string), i[1:]...)
					return false
				}
				return true
			},
		},
	}
	repo := newNodeTypeRepo(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.BatchCreateNodeTypes(tt.args.ctx, tt.args.nodeTypes)
			if !tt.wantErr(t, err, fmt.Sprintf("BatchCreateNodeTypes(%v, %v), err: %v", tt.args.ctx, tt.args.nodeTypes, err)) {
				return
			}
			tt.want(t, tt.args.nodeTypes, got)
		})
	}
}
