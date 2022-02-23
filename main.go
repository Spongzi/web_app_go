package main

import (
	"WebApp/dao/mysql"
	"WebApp/dao/redis"
	"WebApp/logger"
	"WebApp/settings"
	"fmt"
)

func main() {
	// 1. 读取配置
	if err := settings.Init(); err != nil {
		fmt.Println("settings init failed", err)
	}
	// 2. 设置日志
	if err := logger.Init(); err != nil {
		fmt.Println("Logger init failed", err)
	}
	// 3. 连接数据库
	if err := mysql.Init(); err != nil {
		fmt.Println("Mysql init failed", err)
	}
	// 4. 连接redis
	if err := redis.Init(); err != nil {
		fmt.Println("redis init failed", err)
	}
	// 5. 路由管理
}
