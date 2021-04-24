package model

import (
	"github.com/jinzhu/gorm"
	"github.com/tzr2020/gin_blog/utils/errmsg"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// 查询分类是否存在
func CheckCategory(name string) int {
	var c Category
	err = db.Select("id").Where("name=?", name).First(&c).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return errmsg.ERROR
	}
	if c.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCESS
}

// 新增分类
func CreateCategory(data *Category) int {
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询分类列表
func GetCategories(pageSize int, pageNum int) []Category {
	var cs []Category
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cs
}

// 编辑分类
func EditCategory(id int, data *Category) int {
	var c Category
	var m = make(map[string]interface{})
	m["name"] = data.Name
	err = db.Model(&c).Where("id=?", id).Updates(m).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除分类
func DeleteCategory(id int) int {
	var c Category
	err = db.Where("id=?", id).Delete(&c).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询某个分类的文章列表
