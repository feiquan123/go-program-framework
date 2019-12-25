package utils

import (
	"github.com/onrik/logrus/filename"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
)

func NewLogger(logfile, level string, isJson, console, displayline bool) (log *logrus.Logger) {
	log = logrus.New()

	// set level
	switch strings.ToUpper(level) {
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
	if displayline {
		// abs path not use
		//log.SetReportCaller(v.GetBool("log.displayLine"))

		// relative path
		filenameHook := filename.NewHook()
		filenameHook.Field = "file" // custom field
		log.AddHook(filenameHook)
	}

	// log msg format
	if isJson {
		log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05.000",
		})
	} else {
		log.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05.000",
		})
	}

	// output
	if console {
		log.SetOutput(os.Stdout)
	} else {
		var err error
		logfile, err = filepath.Abs(logfile)
		if err != nil {
			log.Fatalln(err)
		}
		// create log fold
		if err = os.MkdirAll(filepath.Dir(logfile), os.ModePerm); err != nil {
			log.Fatalln("can not create log fold,", filepath.Dir(logfile))
		}
		// out put to file
		file, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModeAppend)
		if err != nil {
			log.Fatalln("failed to open log file,", err)
		}
		// modify log file authority
		if err := os.Chmod(logfile, 0666); err != nil {
			log.Fatalln("can not modify log file authority to 0666,", err)
		}
		log.Out = file
	}
	return log
}
