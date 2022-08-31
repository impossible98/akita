package model

import (
	// import third-party packages
	"gorm.io/gorm"
)

type Problem struct {
	gorm.Model
	Identity   string `gorm:"type:varchar(36);column:identity;" json:"identity"`        // 问题表的唯一标识
	CategoryId string `gorm:"type:varchar(255);column:category_id;" json:"category_id"` // 问题所属的分类ID
	Title      string `gorm:"type:varchar(255);column:title;" json:"title"`             // 文章标题
	Content    string `gorm:"type:text;column:content;" json:"content"`                 // 文章内容
	MaxRuntime int    `gorm:"type:int(11);column:max_runtime;" json:"max_runtime"`      // 最大运行时间
	MaxMemory  int    `gorm:"type:int(11);column:max_memory;" json:"max_memory"`        // 最大运行内存
}

func (table Problem) TableName() string {
	return "problem"
}
