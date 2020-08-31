package initialize

import "gin-demo/utils"

func init() {
	Initconfigfile()
	Initlog()
	utils.InitGlobalTrans()
	Mysql()
	MigrateTables()
}
