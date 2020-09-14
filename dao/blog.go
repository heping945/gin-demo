package dao

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title           string `json:"title" gorm:"not null;unique;comment:'标题'"`
	Content         string `json:"content" gorm:"type:longtext;comment:'内容'"`
	MdContent       string `json:"md_content" gorm:"not null;type:longtext;comment:'md内容'"`
	Excerpt         string `json:"excerpt" gorm:"comment:'摘要'"`
	IfReproduce     bool   `json:"if_reproduce" gorm:"default:false;comment:'是否转载'"`
	ReproduceSource string `json:"reproduce_source" gorm:"comment:'转载地址'"`
	CanComment      bool   `json:"can_comment" gorm:"default:true;comment:'是否可以评论'"`
	IfShow          bool   `json:"if_show" gorm:"default:true;comment:'是否展示'"`
	AuthorID        User   `json:"author" gorm:"comment:'作者'"`
	CategoryID      uint   `json:"category_id" `
	Tags            []Tag  `gorm:"many2many:post_tags;" json:"tags"`
}

func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	if p.Excerpt == "" {
		l := len(p.MdContent)
		if l <= 100 {
			p.Excerpt = p.MdContent
		} else {
			p.Excerpt = p.MdContent[:100]
		}
	}

	return
}

type Category struct {
	gorm.Model
	Name             string    `json:"name" gorm:"not null;comment:'分类名'"`
	Slug             string    `json:"slug" gorm:"unique_index;comment:'slug'"`
	ParentCategory   *Category `json:"parent_category" gorm:"foreignKey:ParentCategoryID;comment:'父分类'"`
	ParentCategoryID uint      `json:"parent_category_id" gorm:"comment:'父分类'"`
	Posts            Post      `json:"posts"`
}

type Tag struct {
	gorm.Model
	Name string `json:"name" gorm:"not null;comment:'标签名'"`
	Slug string `json:"slug" gorm:"unique_index;comment:'slug'"`
}
