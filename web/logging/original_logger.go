package logging

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var logger *log.Logger

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

	logger = log.New(multiWriter, "web app | ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
	// set writer.
	logger.SetOutput(multiWriter)
	// output test log.
	logger.Println("set logger !!")
	return file
}

func GetLogger() *log.Logger {
	return logger
}
