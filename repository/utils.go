package repository

import "gorm.io/gorm"

var defaultMaxPaginationSize int = 10

func Paginate(page int) func(db *gorm.DB) *gorm.DB {

	return func(db *gorm.DB) *gorm.DB {
		if page < 1 {
			page = 1
		}
		offset := (page - 1) * defaultMaxPaginationSize
		return db.Offset(offset).Limit(defaultMaxPaginationSize)
	}
}
