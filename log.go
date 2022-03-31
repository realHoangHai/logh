package logh

import (
	"io"
	"log"
	"os"
)

// Level type
type level int

const (
	// DEBUG level. Usually only enabled when debugging. Very verbose logging.
	DEBUG level = iota
	// TRACE level. Designates finer-grained informational events than the Debug.
	TRACE
	// INFO level. General operational entries about what's going on inside the application.
	INFO
	// WARN level. Non-critical entries that deserve eyes.
	WARN
	// ERROR level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ERROR
	// PANIC level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PANIC
	// FATAL level. Logs and then calls `logger.Exit(1)`.
	// It will exit even if the logging level is set to Panic.
	FATAL

	flag = log.Ldate | log.Ltime
)

// Log level prefix map
var prefix = map[level]string{
	DEBUG: "DEBUG: ",
	TRACE: "TRACE: ",
	INFO:  "INFO: ",
	WARN:  "WARN: ",
	ERROR: "ERROR: ",
	PANIC: "PANIC: ",
	FATAL: "FATAL: ",
}

// Loghs ...
type Loghs map[level]LogH

// New returns instance of Logger
func New(out, errOut io.Writer, f Formatter) Loghs {
	// Fall back to stdout if out not set
	if out == nil {
		out = os.Stdout
	}

	// Fall back to stderr if errOut not set
	if errOut == nil {
		errOut = os.Stderr
	}

	// Fall back to DefaultFormatter if f not set
	if f == nil {
		f = new(DefaultFormatter)
	}

	l := make(map[level]LogH, 7)
	l[DEBUG] = &Logger{lvl: DEBUG, formatter: f, logh: log.New(out, f.GetPrefix(DEBUG)+prefix[DEBUG], flag)}
	l[TRACE] = &Logger{lvl: INFO, formatter: f, logh: log.New(out, f.GetPrefix(TRACE)+prefix[TRACE], flag)}
	l[INFO] = &Logger{lvl: INFO, formatter: f, logh: log.New(out, f.GetPrefix(INFO)+prefix[INFO], flag)}
	l[WARN] = &Logger{lvl: INFO, formatter: f, logh: log.New(out, f.GetPrefix(WARN)+prefix[WARN], flag)}
	l[ERROR] = &Logger{lvl: INFO, formatter: f, logh: log.New(errOut, f.GetPrefix(ERROR)+prefix[ERROR], flag)}
	l[PANIC] = &Logger{lvl: INFO, formatter: f, logh: log.New(errOut, f.GetPrefix(PANIC)+prefix[PANIC], flag)}
	l[FATAL] = &Logger{lvl: INFO, formatter: f, logh: log.New(errOut, f.GetPrefix(FATAL)+prefix[FATAL], flag)}
	return Loghs(l)
}

// Logger ...
type Logger struct {
	lvl       level
	formatter Formatter
	logh      LogH
}

// Print ...
func (w *Logger) Print(v ...interface{}) {
	v = w.formatter.Format(w.lvl, v...)
	v = append(v, w.formatter.GetSuffix(w.lvl))
	w.logh.Print(v...)
}

// Printf ...
func (w *Logger) Printf(format string, v ...interface{}) {
	suffix := w.formatter.GetSuffix(w.lvl)
	v = w.formatter.Format(w.lvl, v...)
	w.logh.Printf("%s"+format+suffix, v...)
}

// Fatal ...
func (w *Logger) Fatal(v ...interface{}) {
	v = w.formatter.Format(w.lvl, v...)
	v = append(v, w.formatter.GetSuffix(w.lvl))
	w.logh.Fatal(v...)
}

// Fatalf ...
func (w *Logger) Fatalf(format string, v ...interface{}) {
	suffix := w.formatter.GetSuffix(w.lvl)
	v = w.formatter.Format(w.lvl, v...)
	w.logh.Fatalf("%s"+format+suffix, v...)
}

// Panic ...
func (w *Logger) Panic(v ...interface{}) {
	v = w.formatter.Format(w.lvl, v...)
	v = append(v, w.formatter.GetSuffix(w.lvl))
	w.logh.Fatal(v...)
}

// Panicf ...
func (w *Logger) Panicf(format string, v ...interface{}) {
	suffix := w.formatter.GetSuffix(w.lvl)
	v = w.formatter.Format(w.lvl, v...)
	w.logh.Panicf("%s"+format+suffix, v...)
}
