// Package handler @Author:冯铁城 [17615007230@163.com] 2023-10-12 10:32:10
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"net-project-1/models"
)

// AddGrade 添加年级
func AddGrade(c *gin.Context) {

	//1.声明参数
	var gradeDto models.GradeDto

	//2.获取参数,如果获取异常,返回400
	if err := c.ShouldBindJSON(&gradeDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	//3.保存数据层
	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "操作成功",
		"data":    gradeDto.SaveGrade(),
	})
}

// DeleteGrade 删除年级
func DeleteGrade(c *gin.Context) {

	//1.声明参数
	var gradeDto models.GradeDto

	//2.获取参数,如果获取异常,返回400
	if err := c.ShouldBindUri(&gradeDto.BaseDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	//3.删除数据
	gradeDto.DeleteGrade()

	//4.返回
	c.JSON(http.StatusNoContent, gin.H{})
}

// GetGrade 查询年级
func GetGrade(c *gin.Context) {

	//1.声明参数
	var gradeDto models.GradeDto

	//2.获取参数,如果获取异常,返回400
	if err := c.ShouldBindUri(&gradeDto.BaseDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	//3.查询返回
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "操作成功",
		"data":    gradeDto.GetGrade(),
	})
}

// UpdateGrade 更新年级
func UpdateGrade(c *gin.Context) {

	//1.声明参数
	var gradeDto models.GradeDto

	//2.获取id参数,如果获取异常,返回400
	if err := c.ShouldBindUri(&gradeDto.BaseDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	//3.获取body参数
	if err := c.ShouldBindJSON(&gradeDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	//4.更新
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "操作成功",
		"data":    gradeDto.UpdateGrade(),
	})
}

// PageGrade 分页查询年级
func PageGrade(c *gin.Context) {

	//1.声明参数
	var gradeVo models.GradeVo

	//2.参数绑定
	if err := c.ShouldBindQuery(&gradeVo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	//3.分页查询
	gradeVos, total := gradeVo.PageGrade()
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "操作成功",
		"data": gin.H{
			"pageNum":  gradeVo.PageNum,
			"pageSize": gradeVo.PageSize,
			"total":    total,
			"data":     gradeVos,
		},
	})
}
