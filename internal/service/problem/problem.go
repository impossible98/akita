package problem

import (
	// import local packages
	"fmt"

	"github.com/impossible98/akita/internal/database"
	"github.com/impossible98/akita/internal/model"
)

func GetProblemList() {
	problems := make([]*model.Problem, 0)
	database.Db.Find(&problems)
	for _, problem := range problems {
		fmt.Println("problem ==> %v \n", problem)
	}
}
