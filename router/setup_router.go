package router

import (
	"github.com/20gu00/go_bookManager/controller"
	"github.com/20gu00/go_bookManager/middleware"
	"github.com/gin-gonic/gin"
)

//用于登录和注册的路由
func SetupApiRouters(r *gin.Engine) {
	r.POST("/register", controller.RegisterHandler)
	r.POST("/login", controller.LoginHandler)
	v1 := r.Group("/api/v1")
	r.Use(middleware.AuthMiddleware())
	v1.POST("/book", controller.CreateBookHandler)
	v1.GET("/book/:id", controller.GetBookDetailHandler)
	v1.GET("/books", controller.GetBookListHandler)
	v1.PUT("/book", controller.UpdateBookHandler)
	v1.DELETE("/book/:id", controller.DeleteBookHandler)
}
