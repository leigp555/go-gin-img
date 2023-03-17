package core

import (
	"context"
	"github.com/redis/go-redis/v9"
	"img/server/global"
	"time"
)

// LinkRedisDB 连接redis数据库
func LinkRedisDB() {
	var redisConf = global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr,
		Password: redisConf.Password,
		DB:       redisConf.DB,       // use default DB
		PoolSize: redisConf.PoolSize, // 连接池大小
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.SugarLog.Fatalf("redis数据库连接失败%v\n", err)
	}
	global.Redb = rdb
	global.SugarLog.Info("成功连接redis数据库")
}
