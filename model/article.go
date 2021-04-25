package model

import (
	"github.com/jinzhu/gorm"
	"github.com/tzr2020/gin_blog/utils/errmsg"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

// 新增文章
func CreateArticle(data *Article) int {
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询文章列表
func GetArticles(pageSize int, pageNum int) ([]Article, int, int) {
	var as []Article
	var total int
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&as).Error
	err = db.Model(&as).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return as, errmsg.SUCCESS, total
}

// 查询某个分类的文章列表
func GetArticlesByCategory(id int, pageSize int, pageNum int) ([]Article, int, int) {
	var as []Article
	var total int
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid=?", id).Find(&as).Error
	err = db.Model(&as).Where("cid=?", id).Count(&total).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return as, errmsg.SUCCESS, total
}

// 查询单个文章
func GetArticleInfo(id int) (Article, int) {
	var a Article
	err = db.Preload("Category").Where("id=?", id).First(&a).Error
	if err != nil {
		return a, errmsg.ERROR_ART_NOT_EXIST
	}
	return a, errmsg.SUCCESS
}

// 编辑文章
func EditArticle(id int, data *Article) int {
	var a Article
	var m = make(map[string]interface{})
	m["title"] = data.Title
	m["cid"] = data.Cid
	m["desc"] = data.Desc
	m["content"] = data.Content
	m["img"] = data.Img
	err = db.Model(&a).Where("id=?", id).Updates(m).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除文章
func DeleteArticle(id int) int {
	var a Article
	err = db.Where("id=?", id).Delete(&a).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
