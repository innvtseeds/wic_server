package logger

import (
	"fmt"
	"log"
	"os"
)

// LogLevel type represents different log levels.
type LogLevel int

const (
	// LogLevelInfo represents the info log level.
	LogLevelInfo LogLevel = iota
	// LogLevelWarning represents the warning log level.
	LogLevelWarning
	// LogLevelError represents the error log level.
	LogLevelError
)

// Logger represents a simple logger service.
type Logger struct {
	logger *log.Logger
}

// NewLogger creates a new Logger instance.
func NewLogger() *Logger {
	return &Logger{
		logger: log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// logMessage logs a message with a specified log level and color.
func (l *Logger) logMessage(level LogLevel, color string, messages ...interface{}) {
	switch level {
	case LogLevelInfo:
		l.logger.Print(color, "INFO: ", fmt.Sprint(messages...), "\x1b[0m") // Reset color
	case LogLevelWarning:
		l.logger.Print(color, "WARNING: ", fmt.Sprint(messages...), "\x1b[0m") // Reset color
	case LogLevelError:
		l.logger.Print(color, "ERROR: ", fmt.Sprint(messages...), "\x1b[0m") // Reset color
	}
}

// Info logs an info message.
func (l *Logger) Info(messages ...interface{}) {
	l.logMessage(LogLevelInfo, "\x1b[34m", messages...) // Green color
}

// Warning logs a warning message.
func (l *Logger) Warning(messages ...interface{}) {
	l.logMessage(LogLevelWarning, "\x1b[33m", messages...) // Yellow color
}

// Error logs an error message.
func (l *Logger) Error(messages ...interface{}) {
	l.logMessage(LogLevelError, "\x1b[31m", messages...) // Red color
}
