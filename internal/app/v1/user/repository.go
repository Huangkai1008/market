package user

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type Repository interface {
	ExistUser(condition map[string]interface{}) (bool, error)
	GetUser(condition map[string]interface{}) (User, error)
	CreateUser(user *User) (*User, error)
	GetAddresses(condition interface{}) (Addresses, error)
	CreateAddress(address *Address) (*Address, error)
	UpdateAddress(addressID uint, userID uint, maps map[string]interface{}) (Address, error)
	DeleteAddress(condition interface{}) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

// ExistUser 查询是否存在用户
func (r *repository) ExistUser(condition map[string]interface{}) (exist bool, err error) {
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

// GetUser 查询单个用户
func (r *repository) GetUser(condition map[string]interface{}) (user User, err error) {
	err = r.db.Where(condition).First(&user).Error
	return
}

// CreateUser 创建用户
func (r *repository) CreateUser(user *User) (*User, error) {
	err := r.db.Create(user).Error
	return user, err
}

// GetAddresses 查询用户收货地址
func (r *repository) GetAddresses(condition interface{}) (addresses Addresses, err error) {
	err = r.db.Where(condition).Find(&addresses).Error
	return addresses, err
}

// CreateAddress 创建用户收货地址
func (r *repository) CreateAddress(address *Address) (*Address, error) {

	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 如果设置此为默认地址, 更新其他默认地址为非默认
	if *address.IsDefault == true {
		condition := make(map[string]interface{})
		condition["is_default"] = true
		condition["user_id"] = address.UserID

		var defaultAddress Address

		if !tx.Where(condition).First(&defaultAddress).RecordNotFound() {
			tx.Model(&defaultAddress).Update(map[string]interface{}{"is_default": false})
		}

	}

	if err := tx.Create(&address).Error; err != nil {
		tx.Rollback()
		return &Address{}, err
	}

	return address, tx.Commit().Error
}

// UpdateAddress 更新用户收货地址
func (r *repository) UpdateAddress(addressID uint, userID uint, maps map[string]interface{}) (address Address, err error) {
	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if maps["is_default"] == true {
		condition := make(map[string]interface{})
		condition["is_default"] = true
		condition["user_id"] = userID

		var defaultAddress Address

		if !tx.Where(condition).First(&defaultAddress).RecordNotFound() && defaultAddress.ID != addressID {
			tx.Model(&defaultAddress).Update(map[string]interface{}{"is_default": false})
		}
	}

	if tx.Where("id = ? and user_id = ?", addressID, userID).First(&address).RecordNotFound() {
		tx.Rollback()
		return Address{}, errors.New("不存在的收货地址")
	}

	if err := tx.Model(&address).Update(maps).Error; err != nil {
		tx.Rollback()
		return Address{}, err
	}

	return address, tx.Commit().Error
}

// DeleteAddress 删除用户收货地址
func (r *repository) DeleteAddress(condition interface{}) (err error) {
	err = r.db.Where(condition).Delete(Address{}).Error
	return
}
