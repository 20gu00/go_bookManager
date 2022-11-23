package controller

import (
	"github.com/20gu00/go_bookManager/dao"
	"github.com/20gu00/go_bookManager/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func RegisterHandler(c *gin.Context) {
	p := new(model.User)
	//context type:application/json
	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	tx := dao.DB.Create(p)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": tx.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "注册成功",
		"data": p.Username,
	})

}

func LoginHandler(c *gin.Context) {
	p := new(model.User)
	//context type:application/json
	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	u := model.User{
		Username: p.Username,
		Password: p.Password,
	}

	if row := dao.DB.Where(&u).First(&u); row == nil {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "用户名或密码错误",
		})
		return
	}

	//数据库做的token,比较重
	token := uuid.New().String()
	if tx := dao.DB.Model(&u).Update("token", token); tx.Error != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": tx.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":   "登陆成功",
		"token": token,
	})
}
