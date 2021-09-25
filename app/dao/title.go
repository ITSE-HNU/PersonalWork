// Package dao 数据库操作层
package dao

import (
	"gitee.com/itse/personal-work/app/entity"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// TitleDaoSet 注入 DI
var TitleDaoSet = wire.NewSet(wire.Struct(new(TitleDao), "*"))

// TitleDao titles 表相关的数据库操作
type TitleDao struct {
	DB *gorm.DB
}

// TitleDealParams title 表数据库操作参数结构体
type TitleDealParams struct {
	Content string
	RoleID  int
}

// Query 根据 content 和 role_id 查询题目
func (t *TitleDao) Query(params TitleDealParams) (*[]entity.Title, error) {
	result := new([]entity.Title)
	db := t.DB.Model(&entity.Title{})

	db = db.Where(map[string]interface{}{"content": params.Content, "role_id": params.RoleID})

	err := db.Find(result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Create 根据 content 和 role_id 创建 title 记录
func (t *TitleDao) Create(params TitleDealParams) error {
	db := t.DB.Model(&entity.Title{})
	title := &entity.Title{
		RoleID:  params.RoleID,
		Content: params.Content,
	}
	err := db.Create(title).Error
	if err != nil {
		return err
	}
	return nil
}
