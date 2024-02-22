package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"mapf/app/warehouse/internal/biz"
	"mapf/internal/data"
	dataerrors "mapf/internal/errors"
)

type affixNodeRepo struct {
	*data.Pager
	data   *Data
	logger *log.Helper
}

func NewAffixNodeRepo(data *Data, logger log.Logger) (biz.AffixNodeRepo, error) {
	return &affixNodeRepo{
		data:   data,
		logger: log.NewHelper(logger),
	}, nil
}

// CreateAffixNode 创建用户节点
func (repo *affixNodeRepo) CreateAffixNode(ctx context.Context, affixNode *biz.AffixNode) (*biz.AffixNode, error) {
	err := repo.data.DB(ctx).Create(affixNode).Error
	if err != nil && errors.Is(err, gorm.ErrDuplicatedKey) {
		return nil, dataerrors.NewDuplicatedKeyError("AffixNode Name: [%s] Existed", affixNode.Name)
	}
	return affixNode, err
}

// BatchCreateAffixNode 批量创建用户节点
func (repo *affixNodeRepo) BatchCreateAffixNode(ctx context.Context, affixNodes []*biz.AffixNode) ([]*biz.AffixNode, error) {
	err := repo.data.DB(ctx).Create(affixNodes).Error
	return affixNodes, err
}

// UpdateAffixNode 更新用户节点
func (repo *affixNodeRepo) UpdateAffixNode(ctx context.Context, affixNode *biz.AffixNode) (*biz.AffixNode, error) {
	stmt := repo.data.DB(ctx).Model(&biz.AffixNode{Model: &data.Model{ID: affixNode.ID}})
	result := stmt.Updates(affixNode.GetChanges())
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, dataerrors.NewDuplicatedKeyError("AffixNode Name: [%s] Existed", affixNode.Name)
		}
		return nil, result.Error
	}
	return affixNode, nil
}

// GetAffixNodeById 通过ID获取用户节点
func (repo *affixNodeRepo) GetAffixNodeById(ctx context.Context, id int) (affixNode *biz.AffixNode, err error) {
	err = repo.data.DB(ctx).Where(biz.AffixNode{Model: &data.Model{ID: id}}).First(&affixNode).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return
}

// SelectAffixNodes 查询用户节点，可分页
func (repo *affixNodeRepo) SelectAffixNodes(ctx context.Context, query *biz.AffixNode, page *data.Page) ([]*biz.AffixNode, *data.Page, error) {
	var nodes []*biz.AffixNode
	var err error
	db, page, err := repo.Pager.Prepare(repo.data.DB(ctx), query, page)
	if err != nil {
		return nil, nil, err
	}
	err = db.Find(&nodes).Error
	return nodes, page, err
}
