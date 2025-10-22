package config

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

// LogLevel represents the logging level
type LogLevel int

const (
	// DEBUG level for detailed diagnostic information
	DEBUG LogLevel = iota
	// INFO level for general informational messages
	INFO
	// WARN level for warning conditions
	WARN
	// ERROR level for error conditions
	ERROR
	// FATAL level for critical errors requiring abort
	FATAL
)

var logLevelNames = map[LogLevel]string{
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
	FATAL: "FATAL",
}

// Logger provides structured logging functionality
type Logger struct {
	level      LogLevel
	fileLogger *log.Logger
	console    bool
	logFile    *os.File
}

// NewLogger creates a new logger instance
func NewLogger(level LogLevel, logFilePath string, console bool) (*Logger, error) {
	logger := &Logger{
		level:   level,
		console: console,
	}

	// Open log file if path is provided
	if logFilePath != "" {
		// Ensure directory exists
		dir := filepath.Dir(logFilePath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create log directory: %w", err)
		}

		// Open log file in append mode
		file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return nil, fmt.Errorf("failed to open log file: %w", err)
		}

		logger.logFile = file
		logger.fileLogger = log.New(file, "", 0) // We'll format timestamps ourselves
	}

	return logger, nil
}

// Close closes the log file
func (l *Logger) Close() error {
	if l.logFile != nil {
		return l.logFile.Close()
	}
	return nil
}

// log writes a log message at the specified level
func (l *Logger) log(level LogLevel, component, message string) {
	if level < l.level {
		return // Skip logs below current level
	}

	timestamp := time.Now().UTC().Format(time.RFC3339)
	levelName := logLevelNames[level]
	formattedMsg := fmt.Sprintf("[%s] [%s] [%s] %s", timestamp, levelName, component, message)

	// Write to file if available
	if l.fileLogger != nil {
		l.fileLogger.Println(formattedMsg)
	}

	// Write to console if enabled
	if l.console {
		var output io.Writer = os.Stdout
		if level >= ERROR {
			output = os.Stderr
		}
		fmt.Fprintln(output, formatConsoleMessage(level, component, message))
	}

	// Exit on FATAL
	if level == FATAL {
		if l.logFile != nil {
			l.logFile.Close()
		}
		os.Exit(1)
	}
}

// formatConsoleMessage formats a message for console output with colors/emojis
func formatConsoleMessage(level LogLevel, component, message string) string {
	var prefix string
	switch level {
	case DEBUG:
		prefix = "🔍 [DEBUG]"
	case INFO:
		prefix = "ℹ️  [INFO]"
	case WARN:
		prefix = "⚠️  [WARN]"
	case ERROR:
		prefix = "❌ [ERROR]"
	case FATAL:
		prefix = "💀 [FATAL]"
	}

	if component != "" {
		return fmt.Sprintf("%s [%s] %s", prefix, component, message)
	}
	return fmt.Sprintf("%s %s", prefix, message)
}

// Debug logs a debug message
func (l *Logger) Debug(component, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	l.log(DEBUG, component, message)
}

// Info logs an informational message
func (l *Logger) Info(component, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	l.log(INFO, component, message)
}

// Warn logs a warning message
func (l *Logger) Warn(component, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	l.log(WARN, component, message)
}

// Error logs an error message
func (l *Logger) Error(component, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	l.log(ERROR, component, message)
}

// Fatal logs a fatal error and exits the program
func (l *Logger) Fatal(component, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	l.log(FATAL, component, message)
}

// Success logs a success message (as INFO level with checkmark)
func (l *Logger) Success(component, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	if l.console {
		fmt.Printf("✅ [SUCCESS] [%s] %s\n", component, message)
	}
	l.log(INFO, component, "[SUCCESS] "+message)
}

// SetLevel sets the minimum log level
func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

// GetLevel returns the current log level
func (l *Logger) GetLevel() LogLevel {
	return l.level
}

// Global logger instance
var globalLogger *Logger

// InitGlobalLogger initializes the global logger
func InitGlobalLogger(level LogLevel, logFilePath string, console bool) error {
	logger, err := NewLogger(level, logFilePath, console)
	if err != nil {
		return err
	}
	globalLogger = logger
	return nil
}

// GetGlobalLogger returns the global logger instance
func GetGlobalLogger() *Logger {
	if globalLogger == nil {
		// Create a default logger if none exists
		globalLogger, _ = NewLogger(INFO, "", true)
	}
	return globalLogger
}

// CloseGlobalLogger closes the global logger
func CloseGlobalLogger() error {
	if globalLogger != nil {
		return globalLogger.Close()
	}
	return nil
}
