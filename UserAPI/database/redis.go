package database

import (
	"github.com/gomodule/redigo/redis"
)

var Pool *redis.Pool	// 创建redis连接池

func init(){
	Pool = &redis.Pool{	//实例化一个连接池
		MaxIdle:16,	//最初的连接数量
		MaxActive:0,	//连接池最大连接数量，0表示按需分配
		IdleTimeout:300,	//连接关闭时间300秒（300秒内不使用自动关闭）
		Dial: func() (redis.Conn, error) {	//要连接的数据库
			return redis.Dial("tcp","localhost:6379")
		},
	}
}

