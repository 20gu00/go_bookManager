package middleware

import (
	"github.com/20gu00/go_bookManager/dao"
	"github.com/20gu00/go_bookManager/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		//获取token，根据token查询用户，如果查不到用户，那就token验证失败
		token := c.Request.Header.Get("token")
		u := &model.User{}
		row := dao.DB.Where("token = ?", token).First(&u).RowsAffected
		if row != 1 {
			c.JSON(http.StatusForbidden, gin.H{
				"msg": "token验证失败",
			})
			//跳过后面处理业务的路由，就是不执行后面的路由方法了
			c.Abort()
		}
		c.Set("UserId", u.ID)
	}
}
