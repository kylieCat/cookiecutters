package logger

import (
	"log"
	"os"
)

type Logger struct {
	context []interface{}
	debug *log.Logger
	info *log.Logger
	notice *log.Logger
	warning *log.Logger
	error *log.Logger
	critical *log.Logger
}

const logFormat = log.Lshortfile|log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC

func NewLogger() *Logger {
	return &Logger{
		debug: log.New(os.Stdout, "DEBUG: ", logFormat),
		info: log.New(os.Stdout, "INFO: ", logFormat),
		notice: log.New(os.Stdout, "NOTICE: ", logFormat),
		warning: log.New(os.Stdout, "WARNING: ", logFormat),
		error: log.New(os.Stdout, "ERROR: ", logFormat),
		critical: log.New(os.Stdout, "CRITICAL: ", logFormat),
	}
}

var logger = NewLogger()

func AddContext(value string) *Logger {
	logger.context = append(logger.context, value)
	return logger
}

func (l Logger) Debug(args ...interface{}) {
	args = append(args, l.context...)
	l.debug.Println(args...)
}

func (l Logger) Info(args ...interface{}) {
	args = append(args, l.context...)
	l.info.Println(args...)
}

func (l Logger) Notice(args ...interface{}) {
	args = append(args, l.context...)
	l.notice.Println(args...)
}

func (l Logger) Warning(args ...interface{}) {
	args = append(args, l.context...)
	l.warning.Println(args...)
}

func (l Logger) Error(args ...interface{}) {
	args = append(args, l.context...)
	l.error.Println(args)
}

func (l Logger) Critical(args ...interface{}) {
	args = append(args, l.context...)
	l.critical.Println(args...)
}


func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Notice(args ...interface{}) {
	logger.Notice(args...)
}

func Warning(args ...interface{}) {
	logger.Warning(args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Critical(args ...interface{}) {
	logger.Critical(args...)
}
