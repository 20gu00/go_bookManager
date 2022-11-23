package router

import "github.com/gin-gonic/gin"

func TestRouter(r *gin.Engine) {
	//路由组，定义路由组后，后面在生成的路由，会加上路由组的前缀
	v1 := r.Group("/api/v1")
	// /api/v1/test
	v1.GET("test", TestHandler)
}

func TestHandler(c *gin.Context) {
	c.String(200, "ok")
}
