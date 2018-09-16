package main

import (
	"v1/Gin/UserAPI/model"
)

func main() {
	// 1、加载数据库
	//database.TestConn()
	a := model.SessionService{}
	s := &model.Session{
		SessionID: "20150421120000",
		UserType:  "admin",
		NickName:  "df",
	}
	a.SetSession(s)


	// 2、加载路由

	// 3、启动Web服务


}
