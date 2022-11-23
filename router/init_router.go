package router

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.Default()
	TestRouter(r)
	SetupApiRouters(r)
	return r
}
