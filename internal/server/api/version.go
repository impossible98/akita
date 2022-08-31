package api

import (
	// import third-party packages
	"github.com/gin-gonic/gin"
	// import local packages
	"github.com/impossible98/akita/global"
)

func Version(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"version": global.VERSION,
	})
}
