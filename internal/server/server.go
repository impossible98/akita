package server

import (
	// import local packages
	"github.com/impossible98/akita/internal/server/router"
)

func InitServer() {
	router.Router()
}
