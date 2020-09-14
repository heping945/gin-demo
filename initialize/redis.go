package initialize

import (
	"gin-demo/global"
	"gin-demo/utils"
	"github.com/go-redis/redis"
)

func Redis() {
	addr := utils.EnvDefaultVal("REDIS_HOST", "127.0.0.1")

	redisCfg := global.GVA_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     addr + ":6379",
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.GVA_LOG.Error(err)
	} else {
		global.GVA_LOG.Info("redis connect ping response:", pong)
		global.GVA_REDIS = client
	}
}
