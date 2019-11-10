package models

type User struct {
	// 用户模型
	BaseModel
	Username     string `gorm:"type:varchar(100);unique" json:"username"`
	Email        string `gorm:"type:varchar(128);unique" json:"email"`
	HashPassword string `gorm:"type:varchar(256);not null" json:"-"`
}

func ExistUser(condition map[string]interface{}) (exist bool, err error) {
	/**
	是否存在用户
	*/
	var (
		count int
		user  User
	)

	db.Where(condition).Find(&user).Count(&count)

	if count > 0 {
		exist = true
	} else {
		exist = false
	}
	return

}

func CreateUser(user User) (User, error) {
	/**
	创建用户
	*/
	err := db.Create(&user).Error
	return user, err
}

func GetUser(condition map[string]interface{}) (user User, err error) {
	/**
	查询用户
	*/

	err = db.Where(condition).First(&user).Error
	return

}
