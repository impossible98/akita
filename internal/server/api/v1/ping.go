package v1

import (
	// import third-party packages
	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
