package category

import (
	"goblog/app/models"
	"goblog/pkg/logger"
	"goblog/pkg/model"
	"goblog/pkg/route"
	"goblog/pkg/types"
)

// Category 文章分类
type Category struct {
	models.BaseModel

	Name string `gorm:"type:varchar(255);not null;" valid:"name"`
}

// Create 创建分类
func (category *Category) Create() (err error) {
	if err = model.DB.Create(&category).Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}

// All 获取分类数据
func All() ([]Category, error) {
	var categories []Category
	if err := model.DB.Find(&categories).Error; err != nil {
		return categories, err
	}
	return categories, nil
}

// Link 方法用来生成文章链接
func (c Category) Link() string {
	return route.Name2URL("categories.show", "id", c.GetStringID())
}

// Get 通过 ID 获取分类
func Get(idstr string) (Category, error) {
	var category Category
	id := types.StringToUint64(idstr)
	if err := model.DB.First(&category, id).Error; err != nil {
		return category, err
	}
	return category, nil
}
