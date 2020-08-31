package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `json:"uuid" gorm:"comment:'用户UUID'"`
	Username string    `json:"username" gorm:"comment:'用户登录名'"`
	Password string    `json:"-"  gorm:"comment:'用户登录密码'"`
	Nickname string    `json:"nickname" gorm:"default:'系统用户';comment:'用户昵称'" `
	Avatar   string    `json:"avatar" gorm:"default:'https://zhaoheping.com/media/avatars/heping--4355db/avatar/772e5ca0.jpg';comment:'用户头像'"`
}
