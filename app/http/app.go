package main

import (
	"github.com/davidridwann/wlb-test.git/internal/config"
	"github.com/davidridwann/wlb-test.git/pkg/log"
	"github.com/davidridwann/wlb-test.git/pkg/postgres"
	"github.com/gin-gonic/gin"
	"net/http"
)

func startApp(config config.App) error {
	_, err := postgres.Connect(config.DBConnections)
	if err != nil {
		log.Err().Fatalln("Failed to Initialized postgres DB:", err)
	}
	log.Std().Infoln("Database Connected :", config.DBConnections.DBName)

	router := newRoutes(welcome)
	return startServer(router, config)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
	return
}
