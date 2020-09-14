package req

import "gin-demo/dao"

type CPBlog struct {
	Title           string `json:"title" binding:"required,min=2,max=24"`
	Content         string `json:"content" `
	MdContent       string `json:"md_content" binding:"required,min=2"`
	Excerpt         string `json:"excerpt" `
	IfReproduce     bool   `json:"if_reproduce" `
	ReproduceSource string `json:"reproduce_source" binding:"required_with=IfReproduce,omitempty,url"`
	CanComment      bool   `json:"can_comment" `
	IfShow          bool   `json:"if_show" `
	CategoryID      uint   `json:"category_id" binding:"required" `
	Tags            []uint ` json:"tags" `
}

func (b *CPBlog) CreateOrPut() *dao.Post {
	post := &dao.Post{
		Title:           b.Title,
		Content:         b.Content,
		MdContent:       b.MdContent,
		IfReproduce:     b.IfReproduce,
		ReproduceSource: b.ReproduceSource,
		IfShow:          b.IfShow,
		CategoryID:      b.CategoryID,
	}
	return post
}
