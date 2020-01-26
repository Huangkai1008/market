package gorm

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"market/internal/app/v1/product"
	"market/internal/app/v1/user"
	"market/internal/pkg/config"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Options jwt配置可选项
type Options struct {
	*config.Database
	*config.Gorm
}

// New 创建Gorm DB实例
func New(opts *Options) (*gorm.DB, error) {
	db, err := gorm.Open(opts.DBType, opts.DSN())
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to database ...")
	}

	if err = configure(db, opts); err != nil {
		return nil, errors.Wrap(err, "配置Orm错误")
	}

	if opts.EnableAutoMigrate {
		if err = autoMigrate(db); err != nil {
			return nil, errors.Wrap(err, "自动映射数据表出错")
		}
	}

	return db, err
}

// configure 配置gorm
func configure(db *gorm.DB, opts *Options) error {
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(opts.MaxIdleConnections)
	db.DB().SetMaxOpenConns(opts.MaxOpenConnections)
	return nil
}

// autoMigrate 自动映射数据表
func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&user.User{}, &user.Address{}, &product.Category{}, &product.CategorySpec{}, &product.Spu{}, &product.SpuDetail{}).Error
}
