package router

import (
	// import third-party packages
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// import local packages
	"github.com/impossible98/akita/internal/server/api/v1"
)

func Router() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/ping", v1.Ping)
	router.Run()
}
