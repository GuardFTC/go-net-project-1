// Package base @Author:冯铁城 [17615007230@163.com] 2023-10-12 14:24:40
package base

// BaseDto 基础DTO
type BaseDto struct {
	ID int `json:"id" uri:"id" binding:"omitempty,min=1"`
}
