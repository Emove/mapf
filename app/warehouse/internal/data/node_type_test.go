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
					Code: "RAIL",
					Name: "巷道节点",
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
					Code:      "CHARGE",
					Name:      "充电站",
					IsDefault: data.IsDefaultTrue,
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
						Code: "PASSAGE_TEST",
						Name: "通道节点",
					},
					{
						Code: "WORKSTATION_TEST",
						Name: "工作站",
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

func Test_nodeTypeRepo_GetNodeTypeById(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		args    args
		want    func(testingT assert.TestingT, nodeType *biz.NodeType)
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "normal",
			args: args{
				ctx: context.Background(),
				id:  16,
			},
			want: func(t assert.TestingT, got *biz.NodeType) {
				assert.NotNilf(t, got, "nodeType is nil")
				assert.Equal(t, got.ID, 16, "NodeType id not equal to want")
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
			name: "nil",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			want: func(t assert.TestingT, got *biz.NodeType) {
				assert.Nilf(t, got, "nodeType is not nil")
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
	repo := newNodeTypeRepo(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.GetNodeTypeById(tt.args.ctx, tt.args.id)
			if !tt.wantErr(t, err, fmt.Sprintf("GetNodeTypeById(%v, %v)", tt.args.ctx, tt.args.id)) {
				return
			}
			tt.want(t, got)
		})
	}
}

func Test_nodeTypeRepo_GetNodeTypesByIds(t *testing.T) {
	type args struct {
		ctx context.Context
		ids []int
	}
	ids := []int{1, 16, 17}
	tests := []struct {
		name    string
		args    args
		want    func(t assert.TestingT, got []*biz.NodeType)
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "normal",
			args: args{ctx: context.Background(), ids: ids},
			want: func(t assert.TestingT, got []*biz.NodeType) {
				assert.NotNilf(t, got, "nil NodeType slice")
				assert.Greater(t, len(got), 0, "empty NodeType slice")
				idMap := map[int]struct{}{
					1:  {},
					16: {},
					17: {},
				}
				for _, nodeType := range got {
					if _, ok := idMap[nodeType.ID]; !ok {
						t.Errorf("unexpected NodeType: %v", nodeType)
					}
				}
			},
			wantErr: defaultWantErrFunc,
		},
	}
	repo := newNodeTypeRepo(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.GetNodeTypesByIds(tt.args.ctx, tt.args.ids)
			if !tt.wantErr(t, err, fmt.Sprintf("GetNodeTypesByIds(%v, %v)", tt.args.ctx, tt.args.ids)) {
				return
			}
			tt.want(t, got)
		})
	}
}

func Test_nodeTypeRepo_UpdateNodeTypeById(t *testing.T) {
	type args struct {
		ctx      context.Context
		id       int
		nodeType *biz.NodeType
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "nothing",
			args:    args{ctx: context.Background(), id: 15, nodeType: &biz.NodeType{Code: "global"}},
			want:    false,
			wantErr: defaultWantErrFunc,
		},
		{
			name:    "update_code",
			args:    args{ctx: context.Background(), id: 16, nodeType: &biz.NodeType{Code: "global"}},
			want:    true,
			wantErr: defaultWantErrFunc,
		},
		{
			name:    "update_name",
			args:    args{ctx: context.Background(), id: 16, nodeType: &biz.NodeType{Name: "全局节点"}},
			want:    true,
			wantErr: defaultWantErrFunc,
		},
		{
			name:    "invalid_default_value",
			args:    args{ctx: context.Background(), id: 16, nodeType: &biz.NodeType{IsDefault: 0}},
			want:    false,
			wantErr: defaultWantErrFunc,
		},
		{
			name:    "update_default",
			args:    args{ctx: context.Background(), id: 16, nodeType: &biz.NodeType{IsDefault: data.IsDefaultFalse}},
			want:    true,
			wantErr: defaultWantErrFunc,
		},
		{
			name:    "update_all",
			args:    args{ctx: context.Background(), id: 16, nodeType: &biz.NodeType{Code: "GLOBAL", Name: "全部节点", IsDefault: data.IsDefaultTrue}},
			want:    true,
			wantErr: defaultWantErrFunc,
		},
	}
	repo := newNodeTypeRepo(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.UpdateNodeTypeById(tt.args.ctx, tt.args.id, tt.args.nodeType)
			if !tt.wantErr(t, err, fmt.Sprintf("UpdateNodeTypeById(%v, %v, %v)", tt.args.ctx, tt.args.id, tt.args.nodeType)) {
				return
			}
			assert.Equalf(t, tt.want, got, "UpdateNodeTypeById(%v, %v, %v)", tt.args.ctx, tt.args.id, tt.args.nodeType)
		})
	}
}

func Test_nodeTypeRepo_DeleteNodeTypeById(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "normal",
			args:    args{ctx: context.Background(), id: 17},
			want:    true,
			wantErr: defaultWantErrFunc,
		},
		{
			name:    "default_node_type",
			args:    args{ctx: context.Background(), id: 16},
			want:    false,
			wantErr: defaultWantErrFunc,
		},
	}
	repo := newNodeTypeRepo(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.DeleteNodeTypeById(tt.args.ctx, tt.args.id)
			if !tt.wantErr(t, err, fmt.Sprintf("DeleteNodeTypeById(%v, %v)", tt.args.ctx, tt.args.id)) {
				return
			}
			assert.Equalf(t, tt.want, got, "DeleteNodeTypeById(%v, %v)", tt.args.ctx, tt.args.id)
		})
	}
}

func defaultWantErrFunc(t assert.TestingT, err error, i ...interface{}) bool {
	if err != nil {
		t.Errorf(i[0].(string), i[1:]...)
		return false
	}
	return true
}
