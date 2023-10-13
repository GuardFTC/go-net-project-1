// Package base Package models @Author:冯铁城 [17615007230@163.com] 2023-10-12 10:11:31
package base

import "time"

// BaseModel 基础Model
type BaseModel struct {
	ID         int       `gorm:"type:int(8);primaryKey;autoIncrement;comment:主键ID"`
	CreateTime time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdateTime time.Time `gorm:"type:datetime ON UPDATE CURRENT_TIMESTAMP;default:CURRENT_TIMESTAMP;comment:更新时间"`
}
