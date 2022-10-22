package main

import (
	"fmt"
	"github.com/davidridwann/wlb-test.git/internal/config"
	"github.com/davidridwann/wlb-test.git/pkg/log"
	"github.com/pseidemann/finish"
	"net/http"
)

func startServer(handler http.Handler, config config.App) error {
	fin := &finish.Finisher{
		Log: log.Std(),
	}

	server := &http.Server{Addr: fmt.Sprintf(":%s", config.HTTP.Port), Handler: handler}

	fin.Add(server)

	log.Std().Infoln("starting http server at :", config.HTTP.Port)

	err := server.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Err().Fatalln("server error: ", err)
	}

	fin.Wait()

	return err
}
