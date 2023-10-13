// Package utils @Author:冯铁城 [17615007230@163.com] 2023-10-12 11:15:04
package utils

import (
	"fmt"

	"gorm.io/gorm"
)

// CheckDbResult 校验数据库操作结果
func CheckDbResult(result *gorm.DB) {
	if result.Error != nil {
		panic(fmt.Sprintf("db operate error:[%v]", result.Error))
	}
}
