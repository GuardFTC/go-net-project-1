// Package routers @Author:冯铁城 [17615007230@163.com] 2023-10-12 10:25:14
package routers

import (
	"github.com/gin-gonic/gin"
	"net-project-1/handler"
	"net-project-1/handler/global"
)

// Router1 根服务,第1版本路由
var Router1 = gin.Default()

// 声明第一版基础路由
var v1 = Router1.Group("api/v1")

// InitRoutersV1 InitRouters 第1版本路由初始化
func InitRoutersV1() {

	//1.v1使用全局异常处理
	v1.Use(global.ExceptionHandler)

	//2.声明第一版年级路由
	{
		grade := v1.Group("grades")
		grade.POST("", handler.AddGrade)
		grade.DELETE("/:id", handler.DeleteGrade)
		grade.GET("/:id", handler.GetGrade)
		grade.PUT("/:id", handler.UpdateGrade)
		grade.GET("", handler.PageGrade)
	}

	//3.声明第一版班级路由
	{
		classes := v1.Group("/:gradeId/classes")
		classes.POST("", handler.AddClass)
	}
}
