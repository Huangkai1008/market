package account

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Repository interface {
	Get(condition interface{}) (Addresses, error)
	Create(address *Address) (*Address, error)
	Update(addressID uint, userID uint, maps map[string]interface{}) (Address, error)
	Delete(condition interface{}) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

// GetOne 查询用户收货地址
func (r *repository) Get(condition interface{}) (addresses Addresses, err error) {
	err = r.db.Where(condition).Find(&addresses).Error
	return addresses, err
}

// Create 创建用户收货地址
func (r *repository) Create(address *Address) (*Address, error) {

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

func (r *repository) Update(addressID uint, userID uint, maps map[string]interface{}) (address Address, err error) {
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

func (r *repository) Delete(condition interface{}) (err error) {
	err = r.db.Where(condition).Delete(Address{}).Error
	return
}
