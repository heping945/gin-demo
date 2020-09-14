package initialize

import "gin-demo/utils"

func init() {
	Initconfigfile()        // 配置文件
	Initlog()               // 日志系统
	utils.InitGlobalTrans() // validator的中文翻译
	Redis()                 // redis
	Mysql()                 // mysql
	MigrateTables()         // 数据库迁移
}
