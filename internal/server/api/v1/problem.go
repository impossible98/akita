package v1

import (
	// import third-party packages
	"github.com/gin-gonic/gin"
	// import local packages
	"github.com/impossible98/akita/internal/service/problem"
)

func GetProblemList(ctx *gin.Context) {
	problem.GetProblemList()
	ctx.String(200, "get problem list")
}
