package initialize

import (
	"gin-demo/global"
	"gin-demo/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

//初始化数据库并产生数据库全局变量
func Mysql() {
	admin := global.GVA_CONFIG.Mysql
	mysqlHostPort := utils.EnvDefaultVal("MYSQL_HOST", "127.0.0.1") + ":3306"
	mysqlpath := admin.Username + ":" + admin.Password + "@(" + mysqlHostPort + ")/" + admin.Dbname + "?" + admin.Config
	if db, err := gorm.Open("mysql", mysqlpath); err != nil {
		global.GVA_LOG.Panic("MySQL启动异常", err)
	} else {
		global.GVA_DB = db

		//设置连接池
		//空闲
		global.GVA_DB.DB().SetMaxIdleConns(50)
		//打开
		global.GVA_DB.DB().SetMaxOpenConns(100)
		//超时
		global.GVA_DB.DB().SetConnMaxLifetime(time.Second * 30)
		global.GVA_DB.LogMode(admin.LogMode)
	}
}
