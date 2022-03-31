package logger

import (
	"github.com/realHoangHai/logh"
)

var (
	logger = logh.New(nil, nil, new(logh.ColorFormatter))

	// DEBUG ...
	DEBUG = logger[logh.DEBUG]
	// TRACE ...
	TRACE = logger[logh.TRACE]
	// INFO ...
	INFO = logger[logh.INFO]
	// WARN ...
	WARN = logger[logh.WARN]
	// ERROR ...
	ERROR = logger[logh.ERROR]
	// PANIC ...
	PANIC = logger[logh.PANIC]
	// FATAL ...
	FATAL = logger[logh.FATAL]
)

// Set sets a custom logger for all log levels
func Set(l logh.LogH) {
	DEBUG = l
	TRACE = l
	INFO = l
	WARN = l
	ERROR = l
	PANIC = l
	FATAL = l
}

// SetDebug sets a custom logger for DEBUG level logs
func SetDebug(l logh.LogH) {
	DEBUG = l
}

// SetTrace sets a custom logger for TRACE level logs
func SetTrace(l logh.LogH) {
	TRACE = l
}

// SetInfo sets a custom logger for INFO level logs
func SetInfo(l logh.LogH) {
	INFO = l
}

// SetWarning sets a custom logger for WARNING level logs
func SetWarning(l logh.LogH) {
	WARN = l
}

// SetError sets a custom logger for ERROR level logs
func SetError(l logh.LogH) {
	ERROR = l
}

// SetPanic sets a custom logger for PANIC level logs
func SetPanic(l logh.LogH) {
	PANIC = l
}

// SetFatal sets a custom logger for FATAL level logs
func SetFatal(l logh.LogH) {
	FATAL = l
}
