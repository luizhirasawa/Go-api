package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger() *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, "[Logger] ", log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "debug: ", logger.Flags()),
		info:    log.New(writer, "info: ", logger.Flags()),
		warning: log.New(writer, "warning: ", logger.Flags()),
		err:     log.New(writer, "err: ", logger.Flags()),
		writer:  writer,
	}
}

func (l *Logger) Debug(v ...any) {
	l.debug.Println(v...)
}
func (l *Logger) Info(v ...any) {
	l.info.Println(v...)
}
func (l *Logger) Warning(v ...any) {
	l.warning.Println(v...)
}
func (l *Logger) Error(v ...any) {
	l.err.Println(v...)
}

// formated logs
func (l *Logger) DebugF(format string, v ...any) {
	l.debug.Printf(format, v...)
}
func (l *Logger) InfoF(format string, v ...any) {
	l.info.Printf(format, v...)
}
func (l *Logger) WarningF(format string, v ...any) {
	l.warning.Printf(format, v...)
}
func (l *Logger) ErrorF(format string, v ...any) {
	l.err.Printf(format, v...)
}
