package controller

import (
	"github.com/20gu00/go_bookManager/dao"
	"github.com/20gu00/go_bookManager/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateBook(c *gin.Context) {
	p := new(model.Book)
	//context type:application/json
	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	if tx := dao.DB.Create(&p); tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": tx.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "书记创建成功",
	})

}

func GetBookListHandler(c *gin.Context) {
	books := make([]model.Book, 0)
	if tx := dao.DB.Preload("Users").Find(&books); tx != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": tx.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "查询列表成功",
		"data": books,
	})

}

func GetBookDetailHandler(c *gin.Context) {
	idStr := c.Param("id") //get form
	bookId, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	book := model.Book{
		ID: bookId,
	}

	if tx := dao.DB.Preload("User").Find(&book); tx != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": tx.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "详情查询成功",
		"data": book,
	})
}

func UpdataBookHandler(c *gin.Context) {
	p := new(model.Book)
	//context type:application/json
	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	if tx := dao.DB.Where("id = ?", p.ID).Updates(&p); tx != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": tx.Error,
		})
		return
	}

	if len(p.Users) > 0 {
		if tx := dao.DB.Model(&p).Association("Users").Replace(p.Users); tx.Error() != "" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": tx.Error,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg": "更新书籍成功",
		})
	}
}

func DeleteBookHandler(c *gin.Context) {
	idStr := c.Param("id")
	bookId, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}

	if tx := dao.DB.Select("Users").Delete(&model.Book{
		ID: bookId,
	}); tx != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": tx.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "删除书籍成功",
	})
}
