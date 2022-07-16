package logh

import (
	"fmt"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

// Level type
type Level uint32
type Fields map[string]interface{}

// These are the different logging levels. You can set the logging level to log
// on your instance of logger, obtained with `logrus.New()`.
const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)

const (
	maximumCallerDepth int = 25
	knownLogrusFrames  int = 4
	timeFormat             = "2006-01-02 15:04:05.000"
)

var (
	// Used for caller information initialisation
	callerInitOnce     sync.Once
	logrusPackage      string
	minimumCallerDepth = 1
	loghs              = make(map[string]*Logh)
	loghMutex          sync.RWMutex

	defaultLogh = NewLogh(DebugLevel, "default")
)

// Infof logs a formatted info level log to the console
func Infof(format string, v ...interface{}) { defaultLogh.Infof(format, v...) }

// Tracef logs a formatted debug level log to the console
func Tracef(format string, v ...interface{}) { defaultLogh.Tracef(format, v...) }

// Debugf logs a formatted debug level log to the console
func Debugf(format string, v ...interface{}) { defaultLogh.Debugf(format, v...) }

// Warnf logs a formatted warn level log to the console
func Warnf(format string, v ...interface{}) { defaultLogh.Warnf(format, v...) }

// Errorf logs a formatted error level log to the console
func Errorf(format string, v ...interface{}) { defaultLogh.Errorf(format, v...) }

// Fatalf logs a formatted fatal level log to the console then os.Exit(1)
func Fatalf(format string, v ...interface{}) { defaultLogh.Fatalf(format, v...) }

// Panicf logs a formatted panic level log to the console.
// The panic() function is called, which stops the ordinary flow of a goroutine.
func Panicf(format string, v ...interface{}) { defaultLogh.Panicf(format, v...) }

func Init(level string) {
	defaultLogh.SetLevel(StringToLevel(level))
}

func StringToLevel(level string) Level {
	l := logrus.DebugLevel
	switch level {
	case "trace":
		l = logrus.TraceLevel
	case "debug":
		l = logrus.DebugLevel
	case "info":
		l = logrus.InfoLevel
	case "warn":
		l = logrus.WarnLevel
	case "error":
		l = logrus.ErrorLevel
	}
	return Level(l)
}

type Logh struct {
	logger *logrus.Logger
	level  Level
	prefix string
}

func (lh *Logh) Level() string {
	switch lh.level {
	case PanicLevel:
		return "Panic"
	case FatalLevel:
		return "Fatal"
	case ErrorLevel:
		return "Error"
	case WarnLevel:
		return "Warn"
	case InfoLevel:
		return "Info"
	case DebugLevel:
		return "Debug"
	case TraceLevel:
		return "Trace"
	}
	return "Unkown"
}

func (lh *Logh) Prefix() string {
	return lh.prefix
}

func (lh *Logh) SetLevel(level Level) {
	lh.logger.SetLevel(logrus.Level(level))
}

func (lh *Logh) Panicf(format string, v ...interface{}) {
	lh.logger.Panicf(format, v...)
}

func (lh *Logh) Fatalf(format string, v ...interface{}) {
	lh.logger.Fatalf(format, v...)
}

func (lh *Logh) Errorf(format string, v ...interface{}) {
	lh.logger.Errorf(format, v...)
}

func (lh *Logh) Warnf(format string, v ...interface{}) {
	lh.logger.Warnf(format, v...)
}

func (lh *Logh) Debugf(format string, v ...interface{}) {
	lh.logger.Debugf(format, v...)
}

func (lh *Logh) Tracef(format string, v ...interface{}) {
	lh.logger.Tracef(format, v...)
}

func (lh *Logh) Infof(format string, v ...interface{}) {
	lh.logger.Infof(format, v...)
}

func NewLogh(level Level, prefix string) *Logh {
	loghMutex.RLock()
	if logher, found := loghs[prefix]; found {
		loghMutex.RUnlock()
		return &Logh{
			logger: logher.logger,
			level:  level,
			prefix: prefix,
		}
	}
	loghMutex.RUnlock()
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetReportCaller(true)
	l.SetLevel(logrus.Level(level))
	l.SetFormatter(&TextFormatter{
		Prefix:          prefix,
		FullTimestamp:   true,
		TimestampFormat: timeFormat,
		ForceFormatting: true,
	})

	loghMutex.Lock()
	lh := &Logh{
		logger: l,
		level:  level,
		prefix: prefix,
	}
	loghs[prefix] = lh
	loghMutex.Unlock()
	return lh
}

func NewLoghWithFields(level Level, prefix string, fields Fields) *Logh {
	if logher, found := loghs[prefix]; found {
		return logher
	}
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetReportCaller(true)
	l.SetLevel(logrus.Level(level))
	l.SetFormatter(&TextFormatter{
		Prefix:          prefix,
		Fields:          fields,
		FullTimestamp:   true,
		TimestampFormat: timeFormat,
	})

	lh := &Logh{
		logger: l,
		level:  level,
		prefix: prefix,
	}
	loghs[prefix] = lh
	return lh
}

func SetLoghLevel(prefix string, level Level) error {
	if logher, found := loghs[prefix]; found {
		logher.level = level
		logher.logger.SetLevel(logrus.Level(level))
		return nil
	}
	return fmt.Errorf("logger [%v] not found", prefix)
}

func GetLoghs() map[string]*Logh {
	return loghs
}
