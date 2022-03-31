package logger_test

import (
	"examples/logger"
	"testing"
)

func TestDefaultLogger(t *testing.T) {
	logger.DEBUG.Print("debug mode")
	logger.TRACE.Print("trace mode")
	logger.INFO.Print("info mode")
	logger.WARN.Print("warn mode")
	logger.ERROR.Print("error mode")
	logger.PANIC.Print("panic mode")
	logger.FATAL.Print("fatal mod")
}
