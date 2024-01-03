package tx

import (
	"context"
	"gorm.io/gorm"
)

type contextTxKey struct{}

type Transaction interface {
	InTx(context.Context, func(ctx context.Context) error) error
}

// SetTxToContext 将事务对象设置到Context中
func SetTxToContext(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, &contextTxKey{}, tx)
}

// GetTxFromContext 尝试从Context中获取事务对象
func GetTxFromContext(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(&contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return nil
}
