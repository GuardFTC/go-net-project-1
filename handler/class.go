// Package handler @Author:冯铁城 [17615007230@163.com] 2023-10-13 15:34:22
package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"net-project-1/base"
	"net-project-1/models"
)

// AddClass 添加班级
func AddClass(c *gin.Context) {

	//1.创建变量
	var class models.ClassDto

	//2.获取body参数
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	//3.获取uri参数
	if gradeId, err := strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	} else {
		class.GradeId = gradeId
	}

	//4.保存
	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "操作成功",
		"data":    class.AddClass(),
	})
}

// DeleteClass 删除班级
func DeleteClass(c *gin.Context) {

	//1.创建变量
	var class models.ClassDto

	//2.获取gradeId参数
	if gradeId, err := strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	} else {
		class.GradeId = gradeId
	}

	//3.获取classId
	if id, err := strconv.Atoi(c.Param("classId")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	} else {
		class.BaseDto = new(base.BaseDto)
		class.ID = id
	}

	//4.删除数据
	class.DeleteClass()

	//5.返回
	c.JSON(http.StatusNoContent, gin.H{})
}

// GetClass 查询班级
func GetClass(c *gin.Context) {

	//1.创建变量
	var class models.ClassDto

	//2.获取gradeId参数
	if gradeId, err := strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	} else {
		class.GradeId = gradeId
	}

	//3.获取classId
	if id, err := strconv.Atoi(c.Param("classId")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	} else {
		class.BaseDto = new(base.BaseDto)
		class.ID = id
	}

	//4.查询返回
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "操作成功",
		"data":    class.GetClass(),
	})
}

// UpdateClass 更新班级
func UpdateClass(c *gin.Context) {

	//1.创建变量
	var class models.ClassDto

	//2.获取body参数
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	//3.获取gradeId参数
	if gradeId, err := strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	} else {
		class.GradeId = gradeId
	}

	//4.获取classId
	if id, err := strconv.Atoi(c.Param("classId")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	} else {
		class.BaseDto = new(base.BaseDto)
		class.ID = id
	}

}
