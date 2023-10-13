// Package models @Author:冯铁城 [17615007230@163.com] 2023-10-12 10:18:17
package models

import (
	"net-project-1/base"
)

// Subject 科目
type Subject struct {
	name string `gorm:"type:varchar(64);not null;comment:科目名称"`
	base.BaseModel
}
