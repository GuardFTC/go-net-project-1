// Package models @Author:冯铁城 [17615007230@163.com] 2023-10-12 10:15:19
package models

import (
	"sync"

	"net-project-1/base"
	"net-project-1/utils"
)

// Class 班级
type Class struct {
	GradeId int    `gorm:"type:int(8);not null;comment:年级ID"`
	IndexCn string `gorm:"type:varchar(64);not null;comment:班级中文编码"`
	*base.BaseModel
}

// ClassDto 班级DTO
type ClassDto struct {
	GradeId int    `json:"gradeId" uri:"gradeId" binding:"omitempty,min=1"`
	IndexCn string `json:"indexCn" form:"indexCn" binding:"required,min=1,max=64"`
	*base.BaseDto
	mu sync.RWMutex
}

// AddClass 添加班级
func (c *ClassDto) AddClass() *ClassDto {

	//1.加锁
	c.mu.Lock()
	c.mu.Unlock()

	//2.初始化PO
	class := new(Class)

	//3.PO->DTO
	class.GradeId = c.GradeId
	class.IndexCn = c.IndexCn

	//4.保存
	utils.CheckDbResult(base.DB.Create(&class))

	//5.ID回写
	c.BaseDto = new(base.BaseDto)
	c.ID = class.ID

	//6.返回
	return c
}
