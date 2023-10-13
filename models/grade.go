// Package models @Author:冯铁城 [17615007230@163.com] 2023-10-12 10:10:13
package models

import (
	"fmt"
	"sync"
	"time"

	"net-project-1/base"
	"net-project-1/utils"
)

// Grade 年级
type Grade struct {
	IndexCn string `gorm:"type:varchar(64);not null;comment:年级中文编码"`
	*base.BaseModel
}

// GradeDto 年级DTO
type GradeDto struct {
	IndexCn string `json:"indexCn" form:"indexCn" binding:"required,min=1,max=64"`
	mu      sync.RWMutex
	*base.BaseDto
}

// SaveGrade 保存年级
func (g *GradeDto) SaveGrade() *GradeDto {

	//1.加锁
	g.mu.Lock()
	defer g.mu.Unlock()

	//2.DTO->PO
	grade := Grade{IndexCn: g.IndexCn}

	//3.保存
	utils.CheckDbResult(base.DB.Create(&grade))

	//4.设置ID
	g.BaseDto = new(base.BaseDto)
	g.ID = grade.ID

	//5.返回
	return g
}

// DeleteGrade 删除年级
func (g *GradeDto) DeleteGrade() {

	//1.查询数据是否存在
	var count int64
	base.DB.Model(&Grade{}).Where("id = ?", g.ID).Count(&count)

	//2.存在进行删除
	if count != 0 {
		utils.CheckDbResult(base.DB.Delete(&Grade{}, g.ID))
	}
}

// GetGrade 查询年级
func (g *GradeDto) GetGrade() *GradeDto {

	//1.查询
	grade := new(Grade)
	grade.BaseModel = new(base.BaseModel)
	grade.ID = g.ID
	utils.CheckDbResult(base.DB.First(&grade))

	//2.设置name
	g.IndexCn = grade.IndexCn

	//3.返回
	return g
}

// UpdateGrade 更新年级
func (g *GradeDto) UpdateGrade() *GradeDto {

	//1.加锁
	g.mu.Lock()
	defer g.mu.Unlock()

	//2.查询年级
	grade := new(Grade)
	utils.CheckDbResult(base.DB.Model(&Grade{}).First(&grade, g.ID))

	//3.设置名称
	grade.IndexCn = g.IndexCn
	grade.UpdateTime = time.Now()
	fmt.Println(time.Now().Format(time.DateTime))

	//4.更新
	utils.CheckDbResult(base.DB.Save(&grade))

	//5.返回
	return g
}

// GradeVo 年级VO
type GradeVo struct {
	IndexCn string `json:"indexCn" form:"indexCn" binding:"omitempty,min=1,max=64"`
	base.BaseVo
}

// PageGrade 分页查询年级
func (g *GradeVo) PageGrade() ([]*GradeVo, int64) {

	//1.构造查询条件
	query := base.DB.Model(&Grade{})
	query.Order("update_time desc,id asc")
	if g.IndexCn != "" {
		query.Where("index_cn = ?", g.IndexCn)
	}

	//2.拼接分页条件
	pageQuery := query
	pageQuery.Offset((g.PageNum - 1) * g.PageSize).Limit(g.PageSize)

	//3.查询集合
	var gradeVos []*GradeVo
	utils.CheckDbResult(pageQuery.Find(&gradeVos))

	//4.查询总数
	var count int64
	utils.CheckDbResult(query.Count(&count))

	//5.返回
	return gradeVos, count
}
