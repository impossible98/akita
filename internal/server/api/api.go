package api

import (
	// import third-party packages
	"github.com/gin-gonic/gin"
)

func API(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"version_url":  "/version",
		"problems_url": "/api/v1/problems",
	})
}
