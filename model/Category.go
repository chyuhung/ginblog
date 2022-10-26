package model

import (
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	// 分类结构简单，数据量小，不适用软删除等，仅保留关键属性
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// GetCategoryName 获取分类名称
func GetCategoryName(id int) string {
	var category Category
	db.Select("name").Where("id = ?", id).First(&category)
	return category.Name
}

// CheckCategory 查询分类是否存在
func CheckCategory(name string) int {
	var category Category
	db.Select("id").Where("name = ? ", name).First(&category)
	if category.ID > 0 {
		return errmsg.ERROR_CATEGORY_USED
	}
	return errmsg.SUCCSE
}

// CreateCategory 新增分类
func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// GetCategory 查询分类列表
func GetCategory(pageSize int, pageNum int) []Category {
	var category []Category
	// 偏移量，置为-1表示取消分页功能
	offset := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offset = -1
	}
	err = db.Limit(pageSize).Offset(offset).Find(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return category
}

// EditCategory 编辑分类
func EditCategory(id int, data *Category) int {
	var category Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Model(&category).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// DeleteCategory 删除分类
func DeleteCategory(id int) int {
	var category Category
	err = db.Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
