package log

import (
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"sync"
)

var (
	stdOnce   = sync.Once{}
	errOnce   = sync.Once{}
	stdLogger *logrus.Logger
	errLogger *logrus.Logger
)

// Init logger configuration
func Init(cfg Config) {
	stdOutPath := cfg.Out
	stdErrPath := cfg.Err

	var (
		stdOutFile *os.File
		stdErrFile *os.File
		err        error
	)

	if stdOutPath == "/dev/stdout" {
		stdOutFile = os.Stdout
	} else {
		stdOutFile, err = os.OpenFile(stdOutPath, os.O_APPEND, os.ModeAppend)
		if err != nil {
			log.Fatalln("cannot open log file, err:", err)
		}
	}

	if stdErrPath == "/dev/stderr" {
		stdErrFile = os.Stderr
	} else {
		stdErrFile, err = os.OpenFile(stdErrPath, os.O_APPEND, os.ModeAppend)
		if err != nil {
			log.Fatalln("cannot open log file, err:", err)
		}
	}

	boot(stdOutFile, stdErrFile)
}

// boot set up the logger output target
func boot(stdLoggerOutput io.Writer, errLoggerOutput io.Writer) {
	Std().SetOutput(stdLoggerOutput)
	Err().SetOutput(errLoggerOutput)
}

// Std return a singleton instance for standard logger.
func Std() *logrus.Logger {
	stdOnce.Do(func() {
		stdLogger = logrus.New()
	})

	return stdLogger
}

// Err return a singleton instance for error logger.
func Err() *logrus.Logger {
	errOnce.Do(func() {
		errLogger = logrus.New()
	})

	return errLogger
}
