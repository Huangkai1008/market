package models

type User struct {
	// 用户模型
	Model
	Username     string `gorm:"type:varchar(100);unique" json:"username"`
	Email        string `gorm:"type:varchar(128);not null" json:"email"`
	HashPassword string `gorm:"type:varchar(256);not null" json:"-"`
}

func ExistUser(params map[string]interface{}) bool {
	/**
	是否存在用户
	*/
	var (
		count int
		user  User
	)
	maps := make(map[string]interface{})
	if username, exist := params["username"]; exist {
		maps["username"] = username
	}
	if email, exist := params["email"]; exist {
		maps["email"] = email
	}

	db.Where(maps).Find(&user).Count(&count)

	if count > 0 {
		return true
	} else {
		return false
	}
}
