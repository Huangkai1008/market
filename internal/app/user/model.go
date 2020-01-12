package user

import (
	"market/internal/pkg/database"
	"market/internal/pkg/schema"
)

// UserSchema 用户模型
type User struct {
	database.BaseModel
	Username     string `gorm:"type:varchar(100);unique" json:"username"`
	Email        string `gorm:"type:varchar(128);unique" json:"email"`
	HashPassword string `gorm:"type:varchar(256);not null" json:"-"`
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
