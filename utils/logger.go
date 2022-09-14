package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	logSystem string = "icm"
)

type Fields struct {
	New log.Fields
}

func Logger() log.Logger {
	c := LoadConfig()
	l := *log.New()
	l.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	if c.LogLevel == "" {

	}
	l.SetOutput(os.Stdout)
	// l, _ := log.LevelFromString("info")
	level, _ := log.ParseLevel(c.LogLevel)
	l.SetLevel(level)
	return l
}
