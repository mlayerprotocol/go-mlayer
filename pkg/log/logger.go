package log

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	logSystem string = "ML"
)
type CustomFormatter struct{
	FullTimestamp bool
	ForceColors bool
	Prefix string
}

func (f *CustomFormatter) Format(entry *log.Entry) ([]byte, error) {
    var levelColor int
    switch entry.Level {
    case log.InfoLevel:
        levelColor = 34 // Blue
	case log.DebugLevel:
        levelColor = 32 // Green
    case log.WarnLevel:
        levelColor = 33 // Yellow
    case log.ErrorLevel:
        levelColor = 31 // Red
    default:
        levelColor = 37 // White
    }

    // Custom formatted log with colors
    return []byte(fmt.Sprintf("\x1b[%dm[%s] %s\x1b[0m\n", levelColor, entry.Level, entry.Message)), nil
}

type PrefixFormatter struct {
    Prefix     string
    Formatter log.Formatter
}
func (p *PrefixFormatter) Format(entry *log.Entry) ([]byte, error) {
    // Format the original log entry
    message, err := p.Formatter.Format(entry)
    if err != nil {
        return nil, err
    }

    // Prepend the prefix to the log message
    return []byte(fmt.Sprintf("[%s] %s", p.Prefix, message)), nil
}

var Logger = *log.New()

func Initialize(level string) {

	Logger.SetFormatter(&PrefixFormatter{
		Prefix: "GOML",
		Formatter: &CustomFormatter{
			FullTimestamp: true,
			ForceColors: true,
		},
	})
	if level == "" {
		level = "info"
	}
	Logger.SetOutput(os.Stdout) // load from config file
	logLevel, _ := log.ParseLevel(level)
	Logger.SetLevel(logLevel)
}

// func Logger() log.Logger {
// 	c := LoadConfig()
// 	l := *log.New()
// 	l.SetFormatter(&log.TextFormatter{
// 		FullTimestamp: true,
// 	})
// 	if c.LogLevel == "" {

// 	}
// 	l.SetOutput(os.Stdout)
// 	// l, _ := log.LevelFromString("info")
// 	level, _ := log.ParseLevel(c.LogLevel)
// 	l.SetLevel(level)
// 	return l
// }
