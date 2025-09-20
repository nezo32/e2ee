package config

import (
	"os"

	"github.com/charmbracelet/log"
)

type logConfig struct {
	LogLevel string
}

func initLogConfig() *logConfig {
	return &logConfig{
		LogLevel: os.Getenv("LOG_LEVEL"),
	}
}


func (cfg *logConfig) SetupLogger() {
	level, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		level = log.InfoLevel
	}

	log.SetLevel(level)
}