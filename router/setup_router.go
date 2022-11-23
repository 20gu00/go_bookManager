package router

import (
	"github.com/20gu00/go_bookManager/controller"
	"github.com/20gu00/go_bookManager/middleware"
	"github.com/gin-gonic/gin"
)

func SetupApiRouter(r *gin.Engine) {
	r.POST("/register", controller.RegisterHandler)
	r.POST("/login", controller.LoginHandler)

	v1 := r.Group("/api/v1")
	r.Use(middleware.AuthMiddlerware())
	v1.POST("book", controller.CreateBook)
	v1.GET("book", controller.GetBookDetailHandler)
	v1.GET("books", controller.GetBookListHandler)
	r.PUT("PUT", controller.UpdataBookHandler)
	v1.DELETE("book/:id", controller.DeleteBookHandler)

}
