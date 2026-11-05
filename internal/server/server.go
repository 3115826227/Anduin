package server

import (
	"github.com/gin-gonic/gin"
)

func New(mode string) *gin.Engine {
	gin.SetMode(mode)
	router := gin.New()
	router.Use(gin.Recovery())
	return router
}
