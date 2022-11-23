package middleware

import (
	"github.com/20gu00/go_bookManager/dao"
	"github.com/20gu00/go_bookManager/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

//func (c *gin.Context) == Handler
func AuthMiddlerware() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		u := &model.User{}
		row := dao.DB.Where("token=?", token).First(&u).RowsAffected
		if row != 1 {
			c.JSON(http.StatusForbidden, gin.H{
				"msg": "token验证失败",
			})
			c.Abort() //跳过后边的路由
		}

		c.Set("UserId", u.ID)
	}
}
