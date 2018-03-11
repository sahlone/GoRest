package logger

import (
	"log"
	"os"
	"strconv"
	"sync"
)

type Logger struct {
	*log.Logger
	filename string
}

var logger *Logger
var once sync.Once
var DEBUG = false

func init() {
	stringVal := os.Getenv("DEBUG")
	if stringVal == "" {
		stringVal = "false"
	}
	DEBUG, err := strconv.ParseBool(stringVal)

	if err != nil {
		Error("DEBUG environment variable value is corrupted", err)
		Fatal(err)
	}
	Info("Logger initiated with DEBUG=%b", DEBUG)
	CreateLoggerInstance()
}

//GetInstance creates a singleton instance of the  logger
func CreateLoggerInstance() *Logger {
	once.Do(func() {
		logger = createLogger("application.log")
	})
	return logger
}

//Create a logger instance
func createLogger(fname string) *Logger {
	file, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		Error("Error creating logger", err)
		Fatal(err)
	}
	return &Logger{
		filename: fname,
		Logger:   log.New(file, "", log.Ldate|log.Ltime),
	}
}

func Info(format string, v ...interface{}) {
	if !DEBUG {
		return
	}
	logMessage("Info", format, v)
}

func Error(format string, v ...interface{}) {
	logMessage("Error : ", format, v)
}

func logMessage(prefix string, format string, v ...interface{}) {
	go func() {
		if v != nil {
			logger.Printf(prefix+format+"\n", v)
		} else {
			logger.Println(prefix + format)
		}
	}()
}

func Fatal(v ...interface{}) {
	log.Printf("%v", v)
	logger.Fatal(v)
}
