package config

import (
	"bytes"
	"fmt"
	"github.com/davidridwann/wlb-test.git/pkg/config/loader"
	"github.com/davidridwann/wlb-test.git/pkg/env"
	"github.com/davidridwann/wlb-test.git/pkg/log"
	"github.com/davidridwann/wlb-test.git/pkg/postgres"
	"os"
)

const AppName = "wlb-test"

type App struct {
	Log           log.Config      `yaml:"log"`
	HTTP          HTTPServer      `yaml:"http"`
	DBConnections postgres.Config `yaml:"db_connections"`
}

func getAppConfigPath(configKind string) string {
	filename := fmt.Sprintf("%s.%s.yaml", configKind, env.Current())

	if env.IsDevelopment() {
		return fmt.Sprintf("../../config/app/%s", filename)
	}

	return fmt.Sprintf("/etc/config/%s/%s", AppName, filename)
}

func getFileContents(path string) []byte {
	contents, err := os.ReadFile(path)
	if err != nil {
		log.Err().Fatalln("cannot load config from path: ", path, ", err: ", err)
		return nil
	}

	return contents
}

func LoadApplicationConfig() App {
	configKindMain := "main"

	directory := getAppConfigPath(configKindMain)
	log.Std().Infoln("loading config from: ", directory)

	contents := getFileContents(directory)
	reader := bytes.NewReader(contents)

	configLoader := loader.YAMLLoader[App]{}
	result, err := configLoader.Load(reader)

	if err != nil {
		log.Err().Fatalln("cannot load config from file ", directory, ", err: ", err)
		return App{}
	}

	return result
}
