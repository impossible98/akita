package router

import (
	// import third-party packages
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// import local packages
	"github.com/impossible98/akita/internal/server/api"
	"github.com/impossible98/akita/internal/server/api/v1"
)

func Router() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/api", api.API)
	router.GET("/version", api.API)
	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/problems", v1.GetProblemList)
	}
	router.Run()
}
