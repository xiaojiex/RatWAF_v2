package dao

import (
	"RatWAF_v2/global"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func InitRedis() {
	var ctx = context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456", // 没有密码，默认值
		DB:       0,        // 默认DB 0
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Redis连接失败: %v", err))
	}
	fmt.Println("redis初始化成功")
	global.RDB = rdb
}
