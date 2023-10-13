// Package models @Author:冯铁城 [17615007230@163.com] 2023-10-12 10:08:15
package models

import (
	"net-project-1/base"
)

// Student 学生
type Student struct {
	Name    string `gorm:"type:varchar(64);not null;comment:名称"`
	Email   string `gorm:"type:varchar(64);not null;unique;comment:邮箱"`
	Age     int    `gorm:"type:int(8);not null;comment:年龄"`
	GradeId int    `gorm:"type:int(8);not null;comment:年级ID"`
	ClassId int    `gorm:"type:int(8);not null;comment:班级ID"`
	base.BaseModel
}
