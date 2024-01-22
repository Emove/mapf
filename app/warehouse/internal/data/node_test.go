package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"mapf/app/warehouse/internal/biz"
	"testing"
)

func newNodeRepo(t *testing.T) biz.NodeRepo {
	logger := newLogger()
	data, _, err := NewData(newConfig(), logger)
	assert.Nil(t, err, "new data error")
	repo, err := NewNodeRepo(data, logger)
	assert.Nil(t, err, "new Node repo error")
	return repo
}

func Test_nodeRepo_CreateNode(t *testing.T) {
	type args struct {
		ctx  context.Context
		node *biz.Node
	}
	tests := []struct {
		name    string
		args    args
		want    func(want, got *biz.Node)
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "normal",
			args: args{ctx: context.Background(), node: &biz.Node{WarehouseId: 9, X: 1, Y: 4}},
			want: func(want, got *biz.Node) {
				assert.NotNil(t, got, "got nil Node")
				assert.Equal(t, want.WarehouseId, got.WarehouseId)
				assert.Equal(t, want.X, got.X)
				assert.Equal(t, want.Y, got.Y)
			},
			wantErr: ignoreDuplicatedKeyError,
		},
	}
	repo := newNodeRepo(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.CreateNode(tt.args.ctx, tt.args.node)
			if !tt.wantErr(t, err, fmt.Sprintf("CreateNode(%v, %v)", tt.args.ctx, tt.args.node)) {
				return
			}
			tt.want(tt.args.node, got)
		})
	}
}

func Test_nodeRepo_BatchCreateNode(t *testing.T) {
	type args struct {
		ctx  context.Context
		node []*biz.Node
	}
	tests := []struct {
		name    string
		args    args
		want    func(wants, gots []*biz.Node)
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "normal",
			args: args{
				ctx: context.Background(),
				node: []*biz.Node{
					&biz.Node{WarehouseId: 9, X: 1, Y: 5},
					&biz.Node{WarehouseId: 9, X: 1, Y: 6},
					&biz.Node{WarehouseId: 9, X: 1, Y: 7},
				},
			},
			want: func(wants, gots []*biz.Node) {
				assert.Equal(t, len(wants), len(gots))
			},
			wantErr: ignoreDuplicatedKeyError,
		},
	}
	repo := newNodeRepo(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.BatchCreateNode(tt.args.ctx, tt.args.node)
			if !tt.wantErr(t, err, fmt.Sprintf("CreateNode(%v, %v)", tt.args.ctx, tt.args.node)) {
				return
			}
			tt.want(tt.args.node, got)
		})
	}
}

func ignoreDuplicatedKeyError(t assert.TestingT, err error, i ...interface{}) bool {
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return true
		}
		t.Errorf(i[0].(string), i[1:]...)
		return false
	}
	return true
}

func Test_nodeRepo_GetNodeById(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name     string
		args     args
		wantNode func(got *biz.Node)
	}{
		{
			name: "normal",
			args: args{ctx: context.Background(), id: 3},
			wantNode: func(got *biz.Node) {
				assert.NotNil(t, got)
				assert.Equal(t, got.ID, 3)
			},
		},
		{
			name: "not_found",
			args: args{ctx: context.Background(), id: 1},
			wantNode: func(got *biz.Node) {
				assert.Nil(t, got)
			},
		},
	}
	repo := newNodeRepo(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNode, err := repo.GetNodeById(tt.args.ctx, tt.args.id)
			assert.Nil(t, err)
			tt.wantNode(gotNode)
		})
	}
}

func Test_nodeRepo_GetNodesByIds(t *testing.T) {
	type args struct {
		ctx context.Context
		ids []int
	}
	tests := []struct {
		name      string
		args      args
		wantNodes func([]*biz.Node)
		wantErr   assert.ErrorAssertionFunc
	}{
		{
			name: "normal",
			args: args{ctx: context.Background(), ids: []int{3, 4, 5}},
			wantNodes: func(nodes []*biz.Node) {
				assert.Equal(t, 3, len(nodes))
			},
			wantErr: defaultWantErrFunc,
		},
	}
	repo := newNodeRepo(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNodes, err := repo.GetNodesByIds(tt.args.ctx, tt.args.ids)
			if !tt.wantErr(t, err, fmt.Sprintf("GetNodesByIds(%v, %v)", tt.args.ctx, tt.args.ids)) {
				return
			}
			tt.wantNodes(gotNodes)
		})
	}
}

func Test_nodeRepo_GetNodesByWarehouseId(t *testing.T) {
	type args struct {
		ctx         context.Context
		warehouseId int
	}
	tests := []struct {
		name      string
		args      args
		wantNodes func([]*biz.Node)
		wantErr   assert.ErrorAssertionFunc
	}{
		{
			name: "normal",
			args: args{ctx: context.Background(), warehouseId: 9},
			wantNodes: func(nodes []*biz.Node) {
				for _, node := range nodes {
					assert.Equal(t, node.WarehouseId, 9)
				}
			},
			wantErr: defaultWantErrFunc,
		},
	}
	repo := newNodeRepo(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNodes, err := repo.GetNodesByWarehouseId(tt.args.ctx, tt.args.warehouseId)
			if !tt.wantErr(t, err, fmt.Sprintf("GetNodesByWarehouseId(%v, %v)", tt.args.ctx, tt.args.warehouseId)) {
				return
			}
			tt.wantNodes(gotNodes)
		})
	}
}

func Test_nodeRepo_SelectNodeTupleByCoorRange(t *testing.T) {
	type args struct {
		ctx         context.Context
		warehouseId int
		start       int
		end         int
		auxAix      int
		tupleType   biz.NodeTupleType
	}
	tests := []struct {
		name    string
		args    args
		want    func(biz.NodeTuple)
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "normal_column",
			args: args{
				ctx:         context.Background(),
				warehouseId: 9,
				start:       1,
				end:         7,
				auxAix:      1,
				tupleType:   biz.NodeTupleTypeColumn,
			},
			want: func(tuple biz.NodeTuple) {
				for _, node := range tuple {
					assert.Equal(t, node.WarehouseId, 9)
				}
			},
			wantErr: defaultWantErrFunc,
		},
		{
			name: "normal_row",
			args: args{
				ctx:         context.Background(),
				warehouseId: 9,
				start:       1,
				end:         1,
				auxAix:      5,
				tupleType:   biz.NodeTupleTypeRow,
			},
			want: func(tuple biz.NodeTuple) {
				for _, node := range tuple {
					assert.Equal(t, node.WarehouseId, 9)
				}
			},
			wantErr: defaultWantErrFunc,
		},
	}
	repo := newNodeRepo(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.SelectNodeTupleByCoorRange(tt.args.ctx, tt.args.warehouseId, tt.args.start, tt.args.end, tt.args.auxAix, tt.args.tupleType)
			if !tt.wantErr(t, err, fmt.Sprintf("SelectNodeTupleByCoorRange(%v, %v, %v, %v, %v, %v)", tt.args.ctx, tt.args.warehouseId, tt.args.start, tt.args.end, tt.args.auxAix, tt.args.tupleType)) {
				return
			}
			tt.want(got)
		})
	}
}
