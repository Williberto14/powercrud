package utils

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"time"
)

type LoggerUtil interface {
	Info(fromLocalFileName string, params ...interface{}) string
	Success(fromLocalFileName string, params ...interface{}) string
	Error(fromLocalFileName string, params ...interface{}) string
	Fatal(fromLocalFileName string, params ...interface{})
	Debug(fromLocalFileName string, params ...interface{}) string
	Warn(fromLocalFileName string, params ...interface{}) string
}

type loggerUtil struct {
}

// Colors to be used in the different logger levels
const colorReset = "\033[0m"
const colorCyan = "\033[96m"
const colorRed = "\033[91m"
const colorYellow = "\033[93m"
const colorGreen = "\033[92m"
const colorWhite = "\033[37m"

// format for messages
const msg_format = "Level:[%v] Date:[%v] Origin:[%s:%d] Msg:[%v] "

func NewLogger() LoggerUtil {
	loggerUtilInit := loggerUtil{}
	return &loggerUtilInit
}

// Public methods
// ==================================================================================

// Debug method
func (l *loggerUtil) Debug(msg string, params ...interface{}) string {
	pc, _, line, _ := runtime.Caller(1)
	logResult := fmt.Sprintf(
		msg_format,
		"DEBUG",
		getCurrentDateFormated(),
		filepath.Base(runtime.FuncForPC(pc).Name()),
		line,
		msg,
	)
	log.Printf(string(colorWhite)+logResult, params...)
	log.Print(string(colorReset))
	return logResult
}

// Info method
func (l *loggerUtil) Info(msg string, params ...interface{}) string {
	pc, _, line, _ := runtime.Caller(1)
	logResult := fmt.Sprintf(
		msg_format,
		"INFO",
		getCurrentDateFormated(),
		filepath.Base(runtime.FuncForPC(pc).Name()),
		line,
		msg,
	)
	log.Printf(string(colorCyan)+logResult, params...)
	log.Print(string(colorReset))
	return logResult
}

// Success method
func (l *loggerUtil) Success(msg string, params ...interface{}) string {
	pc, _, line, _ := runtime.Caller(1)
	logResult := fmt.Sprintf(
		msg_format,
		"SUCCESS",
		getCurrentDateFormated(),
		filepath.Base(runtime.FuncForPC(pc).Name()),
		line,
		msg,
	)
	log.Printf(string(colorGreen)+logResult, params...)
	log.Print(string(colorReset))
	return logResult
}

// Warn method
func (l *loggerUtil) Warn(msg string, params ...interface{}) string {
	pc, _, line, _ := runtime.Caller(1)
	logResult := fmt.Sprintf(
		msg_format,
		"WARN",
		getCurrentDateFormated(),
		filepath.Base(runtime.FuncForPC(pc).Name()),
		line,
		msg,
	)
	log.Printf(string(colorYellow)+logResult, params...)
	log.Print(string(colorReset))
	return logResult
}

// Error method
func (l *loggerUtil) Error(msg string, params ...interface{}) string {
	pc, _, line, _ := runtime.Caller(1)
	logResult := fmt.Sprintf(
		msg_format,
		"ERROR",
		getCurrentDateFormated(),
		filepath.Base(runtime.FuncForPC(pc).Name()),
		line,
		msg,
	)
	log.Printf(string(colorRed)+logResult, params...)
	log.Print(string(colorReset))
	return logResult
}

// Fatal method
func (l *loggerUtil) Fatal(msg string, params ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	logResult := fmt.Sprintf(
		msg_format,
		"FATAL",
		getCurrentDateFormated(),
		filepath.Base(runtime.FuncForPC(pc).Name()),
		line,
		msg,
	)
	log.Panicf(string(colorRed)+logResult, params...)
	log.Print(string(colorReset))
}

// Private methods
// ==================================================================================

// getCurrentDateFormated method
func getCurrentDateFormated() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
