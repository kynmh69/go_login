package logging

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var logger *Logger

const (
	DEBUG    = 0
	INFO     = 10
	WARN     = 20
	ERROR    = 30
	CRITICAL = 40
)

type Logger struct {
	log.Logger
	LogLevel int8
}

func (l *Logger) Debug(msg ...any) {
	l.logMsg(DEBUG, "DEBUG", msg...)
}

func (l *Logger) Info(msg ...any) {
	l.logMsg(INFO, "INFO", msg...)
}

func (l *Logger) Warn(msg ...any) {
	l.logMsg(WARN, "WARN", msg...)
}

func (l *Logger) Error(msg ...any) {
	l.logMsg(ERROR, "ERROR", msg...)
}

func (l *Logger) logMsg(levelNum int8, level string, msg ...any) {
	tmpMsg := fmt.Sprintf("%v: %v", level, msg)
	if levelNum >= l.LogLevel {
		l.Println(tmpMsg)
	}

}

func New(w io.Writer, prefix string, flag int) *Logger {
	return &Logger{}
}

func SetLogger() *os.File {
	// date format
	const layout = "2006-01-02"
	dateNow := time.Now()
	dayFormated := dateNow.Format(layout)
	// create file name.
	logFileName := "golang_" + dayFormated + ".log"
	logFilePath := "log/" + logFileName

	fmt.Println("open log file.")

	// open file.
	file, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("file open error.", err)
	}

	// create multi writer.
	// file and stdout
	multiWriter := io.MultiWriter(os.Stdout, file)

	logger = New(multiWriter, "web app | ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
	// set writer.
	logger.SetOutput(multiWriter)
	// output test log.
	logger.Println("set logger !!")
	return file
}

func GetLogger() *Logger {
	return logger
}
