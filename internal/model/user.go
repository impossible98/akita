package model

import (
	// import third-party packages
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Identify string `gorm:"type:varchar(36);column:identity;" json:"identity"` // 用户表的唯一标识
	Name     string `gorm:"type:varchar(100);column:name;" json:"name"`        // 用户名
	Password string `gorm:"type:varchar(32);column:password;" json:"password"` // 用户密码
	Email    string `gorm:"type:varchar(255);column:email;" json:"email"`      // 用户邮
}

func (table *User) TableName() string {
	return "user"
}
