// Package handler @Author:冯铁城 [17615007230@163.com] 2023-10-13 15:34:22
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"net-project-1/models"
)

// AddClass 添加班级
func AddClass(c *gin.Context) {

	//1.创建变量
	var class models.ClassDto

	//3.获取body参数
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	//2.获取uri参数
	if err := c.ShouldBindUri(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	//3.保存
	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "操作成功",
		"data":    class.AddClass(),
	})
}
