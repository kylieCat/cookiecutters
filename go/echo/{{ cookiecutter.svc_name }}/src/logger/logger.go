package logger

import (
	"log"
	"os"
)

type Logger struct {
	context  []interface{}
	debug    *log.Logger
	info     *log.Logger
	notice   *log.Logger
	warning  *log.Logger
	error    *log.Logger
	critical *log.Logger
}

const logFormat = log.Lshortfile | log.Ldate | log.Ltime | log.LUTC

var logger *Logger

func Init() {
	context := make([]interface{}, 2)

	context[0] = "application=xray"
	context[1], _ = os.Hostname()

	logger = &Logger{
		debug:    log.New(os.Stdout, "DEBUG: ", logFormat),
		info:     log.New(os.Stdout, "INFO: ", logFormat),
		notice:   log.New(os.Stdout, "NOTICE: ", logFormat),
		warning:  log.New(os.Stdout, "WARNING: ", logFormat),
		error:    log.New(os.Stdout, "ERROR: ", logFormat),
		critical: log.New(os.Stdout, "CRITICAL: ", logFormat),
		context:  context,
	}
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

func (l Logger) Debugf(template string, args ...interface{}) {
	args = append(args, l.context...)
	l.debug.Printf(template, args...)
}

func (l Logger) Infof(template string, args ...interface{}) {
	args = append(args, l.context...)
	l.info.Printf(template, args...)
}

func (l Logger) Noticef(template string, args ...interface{}) {
	args = append(args, l.context...)
	l.notice.Printf(template, args...)
}

func (l Logger) Warningf(template string, args ...interface{}) {
	args = append(args, l.context...)
	l.warning.Printf(template, args...)
}

func (l Logger) Errorf(template string, args ...interface{}) {
	args = append(args, l.context...)
	l.error.Printf(template, args...)
}

func (l Logger) Criticalf(template string, args ...interface{}) {
	args = append(args, l.context...)
	l.critical.Printf(template, args...)
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

func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

func Noticef(template string, args ...interface{}) {
	logger.Noticef(template, args...)
}

func Warningf(template string, args ...interface{}) {
	logger.Warningf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

func Criticalf(template string, args ...interface{}) {
	logger.Criticalf(template, args...)
}
