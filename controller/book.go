package controller

import (
	"github.com/20gu00/go_bookManager/dao"
	"github.com/20gu00/go_bookManager/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//新增
func CreateBookHandler(c *gin.Context) {
	p := new(model.Book)
	//context type: application/json
	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	//数据库落库
	if tx := dao.DB.Create(&p); tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": tx.Error,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "书籍创建成功",
	})
}

//查看列表
func GetBookListHandler(c *gin.Context) {
	books := make([]*model.Book, 0)
	if tx := dao.DB.Preload("Users").Find(&books); tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": tx.Error,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "查询列表成功",
		"data": books,
	})
}

//查看详情（单条记录）
func GetBookDetailHandler(c *gin.Context) {
	idStr := c.Param("id")
	bookId, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	book := &model.Book{
		ID: bookId,
	}
	//数据库查询
	if tx := dao.DB.Preload("Users").Find(&book); tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": tx.Error,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "查询书籍详情成功",
		"data": book,
	})
}

//修改
func UpdateBookHandler(c *gin.Context) {
	//前端会把书籍修改后的详情信息给到，通过id，去修改所有信息
	p := new(model.Book)
	//context type: application/json
	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	//数据库更新
	if tx := dao.DB.Model(&p).Where("id = ?", p.ID).Updates(&p); tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": tx.Error,
		})
		return
	}
	//以上更新只是针对与book，但是还要更新关联表，也就是user
	if len(p.Users) > 0 {
		dao.DB.Model(&p).Association("Users").Replace(&p.Users)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "更新书籍成功",
	})
}

//删除
func DeleteBookHandler(c *gin.Context) {
	idStr := c.Param("id")
	bookId, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	//数据库删除，删除book时，也要删除关联表中的用户对应关系
	tx := dao.DB.Select("Users").Delete(&model.Book{
		ID: bookId,
	})
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": tx.Error,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "删除书籍成功",
	})
}
