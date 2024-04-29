package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug  *log.Logger
	info   *log.Logger
	warn   *log.Logger
	err    *log.Logger
	writer io.Writer
}

func newLogger() *Logger {
	writer := io.Writer(os.Stdout)

	return &Logger{
		debug:  log.New(writer, "DEBUG: ", log.Ldate|log.Ltime),
		info:   log.New(writer, "INFO: ", log.Ldate|log.Ltime),
		warn:   log.New(writer, "WARN: ", log.Ldate|log.Ltime),
		err:    log.New(writer, "ERROR: ", log.Ldate|log.Ltime),
		writer: writer,
	}
}

func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.warn.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

// fomateds

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warn.Printf(format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
