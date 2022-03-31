package logh

import "fmt"

const (
	// Coloring
	resetSeq = "\033[0m"
	colorSeq = "\033[0;%dm"
)

// Color map
var color = map[level]string{
	DEBUG: fmt.Sprintf(colorSeq, 36), // Cyan
	TRACE: fmt.Sprintf(colorSeq, 34), // Blue
	INFO:  fmt.Sprintf(colorSeq, 32), // Green
	WARN:  fmt.Sprintf(colorSeq, 33), // Yellow
	ERROR: fmt.Sprintf(colorSeq, 31), // Red
	PANIC: fmt.Sprintf(colorSeq, 35), // Magenta
	FATAL: fmt.Sprintf(colorSeq, 37), // White
}

// ColorFormatter colors log messages with ASCI escape codes
// and adds filename and line number before the log message
// See https://en.wikipedia.org/wiki/ANSI_escape_code
type ColorFormatter struct {
}

// GetPrefix returns colour escape code
func (cf *ColorFormatter) GetPrefix(lvl level) string {
	return color[lvl]
}

// GetSuffix returns reset sequence code
func (cf *ColorFormatter) GetSuffix(lvl level) string {
	return resetSeq
}

// Format adds filename and line number before the log message
func (cf *ColorFormatter) Format(lvl level, v ...interface{}) []interface{} {
	return append([]interface{}{header()}, v...)
}
