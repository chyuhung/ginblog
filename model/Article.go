package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Category Category `gorm:"foreignKey:Cid"`
	Cid      int      `gorm:"type:int" json:"cid"`
	Title    string   `gorm:"type:varchar(100);not null" json:"title"`
	Desc     string   `gorm:"type:varchar(200)" json:"desc"`
	Content  string   `gorm:"type:longtext" json:"content"`
	Img      string   `gorm:"type:varchar(100)" json:"img"`
}
