package user

import (
	"goblog/app/models"
	"goblog/pkg/logger"
	"goblog/pkg/model"
)

// User 用户模型
type User struct {
	models.BaseModel

	Name     string `gorm:"column:name;type:varchar(255);not null;unique"`
	Email    string `gorm:"column:email;type:varchar(255);default:NULL;unique;"`
	Password string `gorm:"column:password;type:varchar(255)"`

	// gorm:"-" —— 设置 GORM 在读写时略过此字段
	PasswordConfirm string ` gorm:"-" valid:"password_confirm"`
}

func (user *User) Create() (err error) {
	if err = model.DB.Create(&user).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}
