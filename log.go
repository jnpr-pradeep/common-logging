package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// LogLevel represents the different log levels
type LogLevel int

const (
	DEBUG LogLevel = iota + 1
	INFO
	WARNING
	ERROR
)

// GetLogLevel returns the log level based on the environment variable "LOG_LEVEL"
func GetLogLevel() LogLevel {
	levelStr := strings.ToUpper(os.Getenv("LOG_LEVEL"))
	switch levelStr {
	case "DEBUG":
		return DEBUG
	case "INFO":
		return INFO
	case "WARNING":
		return WARNING
	case "ERROR":
		return ERROR
	default:
		return INFO // default log level
	}
}

// GetLogLevel returns the log level based on the environment variable "LOG_LEVEL"
func GetLogLevelStr(logLevel LogLevel) string {
	switch logLevel {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	default:
		return "INFO" // default log level
	}
}

// Log logs the message with the given log level
func Log(level LogLevel, message string) {
	logLevel := GetLogLevel()
	if level < logLevel {
		return // Skip logging if the log level is lower than the current log level
	}

	log.Println(fmt.Sprintf("[%s] %s", GetLogLevelStr(level), message))
}

func main() {
	Log(DEBUG, "Debug message")
	Log(INFO, "Info message")
	Log(WARNING, "Warning message")
	Log(ERROR, "Error message")
}
