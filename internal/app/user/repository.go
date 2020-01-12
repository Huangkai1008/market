package user

import (
	"github.com/jinzhu/gorm"
)

type Repository interface {
	Exist(condition map[string]interface{}) (bool, error)
	GetOne(condition map[string]interface{}) (User, error)
	Create(user *User) (*User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

// Exist 查询是否存在用户
func (r *repository) Exist(condition map[string]interface{}) (exist bool, err error) {
	var (
		count int
		user  User
	)

	r.db.Where(condition).Find(&user).Count(&count)

	if count > 0 {
		exist = true
	} else {
		exist = false
	}
	return
}

// GetOne 查询单个用户
func (r *repository) GetOne(condition map[string]interface{}) (user User, err error) {
	err = r.db.Where(condition).First(&user).Error
	return
}

// Create 创建用户
func (r *repository) Create(user *User) (*User, error) {
	err := r.db.Create(user).Error
	return user, err
}
