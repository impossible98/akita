package model

import (
	// import third-party packages
	"gorm.io/gorm"
)

type Submit struct {
	gorm.Model
	Identity  string `gorm:"type:varchar(36);column:identity;" json:"identity"`     // 提交表的唯一标识
	ProblemId string `gorm:"type:varchar(36);column:problem_id;" json:"problem_id"` // 提交所属的问题ID
	UserId    string `gorm:"type:varchar(36);column:user_id;" json:"user_id"`       // 提交所属的用户ID
	Path      string `gorm:"type:varchar(255);column:path;" json:"path"`            // 提交的路径
}

func (table *Submit) TableName() string {
	return "submit"
}
