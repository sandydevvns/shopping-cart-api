package models

import (
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID          int          `gorm:"primaryKey;autoIncrement"`
	Name        string       `gorm:"type:varchar(191);not null"`
	Description string       `gorm:"default:null"`
	ParentId    int          `gorm:"default:0"`
	Image       string       `gorm:"type:varchar(250);default:null"`
	Slug        string       `gorm:"type:varchar(191);uniqueIndex"`
	Images      []FileUpload `gorm:"foreignKey:CategoryId"`
	IsNewRecord bool         `gorm:"-;default:false"`
	Parent      *Category    `gorm:"-"`
}

func (cat *Category) BeforeSave(*gorm.DB) (err error) {
	cat.Slug = slug.Make(cat.Name)
	return
}

type CatagoryList struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ParentId    int    `json:"parent_id"`
	Filename    string `json:"filename"`
	FilePath    string `json:"file_path"`
	Slug        string `json:"slug"`
	Parent      string `json:"parent"`
}

func (CatagoryList) TableName() string {
	return "categories"
}
