// Package base @Author:冯铁城 [17615007230@163.com] 2023-10-13 09:48:52
package base

// BaseVo 基础VO
type BaseVo struct {
	ID         int      `json:"id"`
	PageNum    int      `json:"-" form:"pageNum" binding:"required,min=1,max=10000" gorm:"-"`
	PageSize   int      `json:"-" form:"pageSize" binding:"required,min=1,max=100" gorm:"-"`
	CreateTime jsonTime `json:"createTime"`
	UpdateTime jsonTime `json:"updateTime"`
}
