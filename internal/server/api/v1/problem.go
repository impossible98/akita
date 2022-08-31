package v1

import (
	// import third-party packages
	"github.com/gin-gonic/gin"
	// import local packages
	"github.com/impossible98/akita/internal/service/problem"
)

func GetProblemList(ctx *gin.Context) {
	result, err := problem.GetProblemList()
	if err != nil {
		ctx.JSON(200, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(200, result)
}
