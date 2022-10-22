package main

import (
	"github.com/gin-gonic/gin"
)

func newRoutes(handler gin.HandlerFunc) *gin.Engine {
	router := gin.Default()

	router.GET("/", handler)

	return router
}
