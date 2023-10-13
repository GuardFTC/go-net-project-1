// Package models @Author:冯铁城 [17615007230@163.com] 2023-10-12 10:21:56
package models

import (
	"net-project-1/base"
)

// Teacher 教师
type Teacher struct {
	Name      string `gorm:"type:varchar(64);not null;comment:名称"`
	Email     string `gorm:"type:varchar(64);not null;unique;comment:邮箱"`
	Age       int    `gorm:"type:int(8);not null;comment:年龄"`
	SubjectId int    `gorm:"type:int(8);not null;comment:教学科目ID"`
	base.BaseModel
}
