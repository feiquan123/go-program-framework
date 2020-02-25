package utils

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/onrik/logrus/filename"
	"github.com/sirupsen/logrus"
)

// LoggerConfig : logger config from yaml or json file
type LoggerConfig struct {
	Level       string
	DisplayLine bool
	IsJSON      bool
	Console     bool
	File        string
}

// NewLoggerConfigFromMap : new logger config from map
func NewLoggerConfigFromMap(l map[string]string) *LoggerConfig {
	log := new(LoggerConfig)
	var ok bool
	if log.Level, ok = l["level"]; !ok {
		panic("can not get [level] form map")
	}

	if log.File, ok = l["file"]; !ok {
		panic("can not get [file] form map")
	}

	if t, ok := l["displayline"]; !ok {
		panic("can not get [displayline] form map")
	} else {
		log.DisplayLine = isTrueStr(t)
	}

	if t, ok := l["isjson"]; !ok {
		panic("can not get [IsJSON] form map")
	} else {
		log.IsJSON = isTrueStr(t)
	}

	if t, ok := l["console"]; !ok {
		panic("can not get [console] form map")
	} else {
		log.Console = isTrueStr(t)
	}
	return log
}

// NewLogger : create a Loggre
func NewLogger(l *LoggerConfig) (log *logrus.Logger) {
	log = logrus.New()

	// set level
	switch strings.ToUpper(l.Level) {
	case "TRACE":
		log.SetLevel(logrus.TraceLevel)
	case "DEBUG":
		log.SetLevel(logrus.DebugLevel)
	case "INFO":
		log.SetLevel(logrus.InfoLevel)
	case "WARNING":
		log.SetLevel(logrus.WarnLevel)
	case "ERROR":
		log.SetLevel(logrus.ErrorLevel)
	case "FATAL":
		log.SetLevel(logrus.FatalLevel)
	case "PANIC":
		log.SetLevel(logrus.PanicLevel)
	}

	// set display line
	if l.DisplayLine {
		// abs path not use
		//log.SetReportCaller(v.GetBool("log.displayLine"))

		// relative path
		filenameHook := filename.NewHook()
		filenameHook.Field = "file" // custom field
		log.AddHook(filenameHook)
	}

	// log msg format
	if l.IsJSON {
		log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05.000",
		})
	} else {
		log.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05.000",
		})
	}

	// output
	if l.Console {
		log.SetOutput(os.Stdout)
	} else {
		var err error
		l.File, err = filepath.Abs(l.File)
		if err != nil {
			log.Fatalln(err)
		}
		// create log fold
		if err = os.MkdirAll(filepath.Dir(l.File), os.ModePerm); err != nil {
			log.Fatalln("can not create log fold,", filepath.Dir(l.File))
		}
		// out put to file
		file, err := os.OpenFile(l.File, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModeAppend)
		if err != nil {
			log.Fatalln("failed to open log file,", err)
		}
		// modify log file authority
		if err := os.Chmod(l.File, 0666); err != nil {
			log.Fatalln("can not modify log file authority to 0666,", err)
		}
		log.Out = file
	}
	return log
}

// NewLoggerFromMap : get log config map from viper
func NewLoggerFromMap(l map[string]string) (log *logrus.Logger) {
	return NewLogger(NewLoggerConfigFromMap(l))
}

// isTrueStr : "true" -> true
func isTrueStr(b string) bool {
	switch strings.ToLower(b) {
	case "true":
		return true
	case "false":
		return false
	default:
		return false
	}
}
