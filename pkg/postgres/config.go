package postgres

import (
	"gorm.io/gorm/logger"
	"time"
)

type Config struct {
	Driver         string          `yaml:"driver"`
	Host           string          `yaml:"host"`
	Port           int             `yaml:"port"`
	User           string          `yaml:"user"`
	Password       string          `yaml:"password"`
	DBName         string          `yaml:"dbname"`
	DBLogEnable    bool            `yaml:"enable"`
	DBLogLevel     logger.LogLevel `yaml:"level"`
	DBLogThreshold time.Duration   `yaml:"threshold"`
}
