package scope

import (
	"math"

	"gorm.io/gorm"
)

type Pagination struct {
	Limit      int
	Page       int
	TotalRows  int64
	TotalPages int
	Rows       interface{}
}

type PaginationQuery struct {
	Limit int `query:"limit"`
	Page  int `query:"page"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func PaginateScope(pagination *Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		session := db.Session(&gorm.Session{})
		var totalRows int64
		session.Count(&totalRows)

		pagination.TotalRows = totalRows
		totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))
		pagination.TotalPages = totalPages
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit())
	}
}
