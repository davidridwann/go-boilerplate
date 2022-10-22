package main

import (
	"github.com/davidridwann/wlb-test.git/internal/config"
	userEntity "github.com/davidridwann/wlb-test.git/internal/entity/user"
	"github.com/davidridwann/wlb-test.git/pkg/log"
	"github.com/davidridwann/wlb-test.git/pkg/mongo"
	"github.com/davidridwann/wlb-test.git/pkg/postgres"
)

func startApp(config config.App) error {
	mongo.Connect()
	db, err := postgres.Connect(config.DBConnections)
	if err != nil {
		log.Err().Fatalln("Failed to Initialized postgres DB:", err)
	}

	err = db.AutoMigrate(&userEntity.User{})
	if err != nil {
		log.Err().Fatalln("Failed to AutoMigrate Entity:", err)
	}

	router := newRoutes(db)
	return startServer(router, config)
}
