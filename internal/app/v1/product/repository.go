package product

import "github.com/jinzhu/gorm"

type Repository interface {
	GetCategory(condition interface{}) (Categories, error)
	GetCategoryCount(condition interface{}) (int, error)
	GetCategorySpec(condition interface{}) (CategorySpecs, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

// GetCategory 获取商品分类
func (r *repository) GetCategory(condition interface{}) (categories Categories, err error) {
	err = r.db.Where(condition).Find(&categories).Error
	return
}

// GetCategoryCount 获取商品分类总数
func (r *repository) GetCategoryCount(condition interface{}) (count int, err error) {
	err = r.db.Model(Category{}).Where(condition).Count(&count).Error
	return
}

// GetCategorySpec 获取商品分类规格信息
func (r *repository) GetCategorySpec(condition interface{}) (specs CategorySpecs, err error) {
	err = r.db.Where(condition).Find(&specs).Error
	return
}
