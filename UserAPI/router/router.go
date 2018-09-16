package router

import (
	"github.com/gin-gonic/gin"
	"v1/Gin/UserAPI/controller"
)

func init(){
	router := gin.Default()
	router.GET("/",controller.IndexApi)
}
