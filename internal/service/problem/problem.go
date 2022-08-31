package problem

import (
	// import local packages
	"github.com/impossible98/akita/internal/database"
	"github.com/impossible98/akita/internal/model"
)

func GetProblemList() (*model.Problem, error) {
	problems := model.Problem{}
	err := database.Db.Find(&problems).Error
	if err != nil {
		return nil, err
	}
	return &problems, nil
}
