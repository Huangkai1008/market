package user

import (
	"market/internal/pkg/database/model"
	"market/internal/pkg/schema"
)

// User 用户模型
type User struct {
	model.BaseModel
	Username     string `gorm:"type:varchar(100);unique;comment:'用户名'" json:"username"` // 用户名
	Email        string `gorm:"type:varchar(128);unique;comment:'邮箱'" json:"email"`     // 邮箱
	HashPassword string `gorm:"type:varchar(256);not null;comment:'密码'" json:"-"`       // 密码
}

func (user *User) ToSchemaUser() (schemaUser *ReadSchema) {
	schemaUser = &ReadSchema{
		BaseSchema: schema.BaseSchema{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt},
		Username: user.Username,
		Email:    user.Email,
	}
	return
}
