package main

import (
	"github.com/davidridwann/wlb-test.git/internal/config"
	"github.com/davidridwann/wlb-test.git/pkg/log"
)

func main() {
	appConfig := config.LoadApplicationConfig()

	err := startApp(appConfig)
	if err != nil {
		log.Err().Fatalf("failed to start app: %v", err)
	}
}
