package main

import (
	"fmt"
	"gin-demo/global"
	"gin-demo/initialize"
	"github.com/gin-gonic/gin"
)

// 引入initialize,init了配置包和日志

func setEnv() {
	Env := global.GVA_CONFIG.System.Env
	gin.SetMode(Env)
	global.GVA_LOG.Debugf("当前是在%s环境", Env)
}

func main() {
	// 程序结束前关闭数据库链接
	defer global.GVA_DB.Close()

	setEnv()
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	global.GVA_LOG.Debugf("本机运行地址：http://127.0.0.1:%d", global.GVA_CONFIG.System.Addr)

	router := initialize.InitRouter()

	router.Run(address)

}
