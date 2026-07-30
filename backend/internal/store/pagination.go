package store

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PaginatedQuery struct {
	Desc    bool
	OrderBy string
}

func (p *PaginatedQuery) paginateResult(db *gorm.DB) *gorm.DB {
	return db.Order(clause.OrderByColumn{
		Column: clause.Column{Name: p.OrderBy}, Desc: p.Desc},
	)
}
