package utils

import (
	"github.com/ipfs/go-log/v2"
)

const (
	logSystem string = "icm"
)

func Logger() *log.ZapEventLogger {
	c := LoadConfig()
	if c.LogLevel == "" {

	}
	l, _ := log.LevelFromString(c.LogLevel)
	log.SetAllLoggers(l)
	return log.Logger(logSystem)
}
