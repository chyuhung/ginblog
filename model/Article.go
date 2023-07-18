package model

import (
	"ginblog/utils/errmsg"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Category Category `gorm:"foreignKey:Cid"`
	Cid      int      `gorm:"type:int" json:"cid"`
	Title    string   `gorm:"type:varchar(100);not null" json:"title"`
	Desc     string   `gorm:"type:varchar(200)" json:"desc"`
	Content  string   `gorm:"type:longtext" json:"content"`
	Img      string   `gorm:"type:varchar(100)" json:"img"`
}

// 迁移数据库自动命名单数形式
func (Article) TableName() string {
	return "article"
}

// GetArticleCid 获取文章cid
func GetArticleCid(id int) int {
	var article Article
	db.Select("cid").Where("id = ?", id).First(&article)
	return article.Cid
}

// CreateArticle 新增文章
func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// GetCategoryArticle 查询分类下所有文章
func GetCategoryArticle(id int, pageSize int, pageNum int) ([]Article, int, int64) {
	var categoryArticleList []Article
	var total int64
	// 偏移量，置为-1表示取消分页功能
	offset := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offset = -1
	}
	err = db.Preload("Category").Limit(pageSize).Offset(offset).Where("cid = ?", id).Find(&categoryArticleList).Count(&total).Error
	if err != nil {
		return categoryArticleList, errmsg.ERROR_CATEGORY_NOT_EXIST, 0
	}
	return categoryArticleList, errmsg.SUCCSE, total
}

// GetArticleInfo 查询单个文章
func GetArticleInfo(id int) (Article, int) {
	var article Article
	err = db.Preload("Category").Where("id = ?", id).First(&article).Error
	if err != nil {
		return article, errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	return article, errmsg.SUCCSE
}

// GetArticle 查询文章列表
func GetArticle(pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var total int64
	// 偏移量，置为-1表示取消分页功能
	offset := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offset = -1
	}
	err = db.Preload("Category").Limit(pageSize).Offset(offset).Find(&articleList).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCSE, total
}

// EditArticle 编辑文章
func EditArticle(id int, data *Article) int {
	var article Article
	var maps = make(map[string]interface{})
	maps["cid"] = data.Cid
	maps["title"] = data.Title
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err = db.Model(&article).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// DeleteArticle 删除文章
func DeleteArticle(id int) int {
	var article Article
	err = db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
