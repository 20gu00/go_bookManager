package router

import "github.com/gin-gonic/gin"

func TestRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	v1.GET("test", TestHandler)
}

func TestHandler(c *gin.Context) {
	c.String(200, "ok")
}
