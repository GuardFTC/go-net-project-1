// Package global @Author:冯铁城 [17615007230@163.com] 2023-10-13 14:25:28
package global

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// ExceptionHandler 全局异常捕获中间件
func ExceptionHandler(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {

			//1.定义异常码
			errorCode := http.StatusInternalServerError

			//2.err转string
			e, ok := r.(string)
			if ok {

				//3.判定是否为数据库唯一索引重复
				if strings.Contains(e, "Duplicate entry") {
					errorCode = http.StatusBadRequest
				} else if strings.Contains(e, "record not found") {
					errorCode = http.StatusNotFound
				}
			}

			//4.返回
			c.JSON(errorCode, gin.H{
				"code":    errorCode,
				"message": e,
			})
		}
	}()

	//5.继续执行
	c.Next()
}
