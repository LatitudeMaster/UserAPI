package main

import (
	"github.com/gin-gonic/gin"
	"database/sql"
	"log"
	"net/http"
	_"github.com/go-sql-driver/mysql" // 这里很重要，导入自己本地使用的数据库驱动，前面是下划线，否则会报错：sql: unknown driver "mysql" (forgotten import?)
	"fmt"
)

func main(){
	db,err := sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/test?parseTime=true")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(20)
	if err := db.Ping();err != nil {
		log.Fatalln(err)
		fmt.Println("db conect erro")
	}

	router := gin.Default()
	router.GET("/",func(c *gin.Context){
		c.String(http.StatusOK,"It works")
	})
	//====================================以上是固定写法
	router.POST("/person", func(c *gin.Context) {
		firstName := c.Request.FormValue("first_name")
		lastName := c.Request.FormValue("last_name")

		rs,err := db.Exec("INSERT INTO person(first_name, last_name) VALUE (?,?)",firstName,lastName)
		if err != nil {
			log.Fatalln(err)
			fmt.Println("insert erro")
		}

		id,err := rs.LastInsertId()
		if err != nil {
			log.Fatalln(err)
			fmt.Println("insert erro")
		}
		fmt.Println("insert person Id {}",id)
		msg := fmt.Sprintf("insert successful %d",id)
		c.JSON(http.StatusOK,gin.H{
			"msg":msg,
		})
	})
//curl -X POST http://127.0.0.1:8000/person -d "first_name=hello&last_name=world" | python -m json.tool

	//====================================以下是固定写法


	router.Run(":8000")
}
