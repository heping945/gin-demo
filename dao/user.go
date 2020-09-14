package dao

import (
	"gin-demo/utils"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `json:"uuid" gorm:"comment:'用户UUID'"`
	Username string    `json:"username" gorm:"not null;unique_index;comment:'用户登录名'"`
	Password string    `json:"-"  gorm:"not null;comment:'用户登录密码'"`
	Nickname string    `json:"nickname" gorm:"default:'系统用户';comment:'用户昵称'" `
	Avatar   string    `json:"avatar" gorm:"default:'https://zhaoheping.com/media/avatars/heping--4355db/avatar/772e5ca0.jpg';comment:'用户头像'"`
	Posts    []Post    `json:"posts" gorm:"foreignKey:AuthorID;comment:'posts'"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = utils.MD5V([]byte(u.Password))
	u.UUID = uuid.NewV4()
	return
}
