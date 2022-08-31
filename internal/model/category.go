package model

import (
	// import third-party packages
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Identity string `gorm:"type:varchar(36);column:identity;" json:"identity"` // 分类表的唯一标识
	Name     string `gorm:"type:varchar(255);column:name;" json:"name"`        // 分类名称
	ParentId int    `gorm:"type:int;column:parent_id;" json:"parent_id"`       // 父分类ID
}

func (table *Category) TableName() string {
	return "category"
}
