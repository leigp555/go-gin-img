package test

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
	"time"
)

// LinkRedisDB 连接redis数据库
func TestRedis(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "1.117.141.66:6379",
		Password: "123456abc",
		DB:       0,  // use default DB
		PoolSize: 50, // 连接池大小
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("redis数据库连接失败%v\n", err)
		return
	}
	fmt.Println("成功连接redis数据库")

	if err := rdb.Set(ctx, "test", "测试", 300*time.Second).Err(); err != nil {
		fmt.Println("数据存储失败")
		return
	}
	val, err := rdb.Get(ctx, "test").Result()
	if err != nil {
		fmt.Println("数据获取失败")
		return
	}
	fmt.Println("获取的数据是：" + val)
}
