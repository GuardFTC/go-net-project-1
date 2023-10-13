// Package models @Author:冯铁城 [17615007230@163.com] 2023-10-12 10:15:19
package models

import (
	"net-project-1/base"
)

// Class 班级
type Class struct {
	GradeId int    `gorm:"type:int(8);not null;comment:年级ID"`
	IndexCn string `gorm:"type:varchar(64);not null;comment:班级中文编码"`
	base.BaseModel
}
