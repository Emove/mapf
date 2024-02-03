package data

import (
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Page 数据查询分页对象
type Page struct {
	Num   int64 `json:"page"`
	Size  int64 `json:"size"`
	Total int64 `json:"total"`
}

type Pager struct {
}

func (*Pager) Prepare(db *gorm.DB, query schema.Tabler, page *Page) (*gorm.DB, *Page, error) {
	if page != nil {
		var total int64
		err := db.Model(query).Where(query).Count(&total).Error
		if err != nil {
			return nil, nil, err
		}
		if page.Num > (total/page.Size)+1 {
			page.Num = (total / page.Size) + 1
		}
		offset := page.Size * (page.Num - 1)
		return db.Where(query).Limit(int(page.Size)).Offset(int(offset)), page, nil
	}
	return db.Where(query), nil, nil
}
