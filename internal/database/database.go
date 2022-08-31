package database

import (
	// import third-party packages
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// import local packages
	"github.com/impossible98/akita/internal/model"
)

var Db = InitDatabase()

func InitDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Category{}, model.Problem{}, model.User{}, model.Submit{})
	return db
}
