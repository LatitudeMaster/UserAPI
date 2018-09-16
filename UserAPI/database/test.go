package database

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func TestConn(){
	c := Pool.Get()	//从连接池，取一个连接
	defer c.Close()	//函数运行结束，把连接放回连接池

	_,err := c.Do("Set","abc",200)
	if err != nil {
		fmt.Println(err)
		return
	}

	r,err := redis.Int(c.Do("Get","abc"))
	if err != nil {
		fmt.Println("get abc failed :",err)
		return
	}
	fmt.Println(r)
	Pool.Close()	//关闭连接池
}
