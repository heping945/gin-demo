package initialize

import (
	"gin-demo/dao"
	"gin-demo/global"
)

func MigrateTables() {
	db := global.GVA_DB
	db.SingularTable(true) //禁用复数
	db.Debug().AutoMigrate(
		dao.User{},
	)
	global.GVA_LOG.Debug("register table success")
}
