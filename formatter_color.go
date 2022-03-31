package logh

import "fmt"

const (
	// Coloring
	resetSeq = "\033[0m"
	colorSeq = "\033[0;%dm"
)

// Color map
var color = map[level]string{
	INFO:  fmt.Sprintf(colorSeq, 32), // green
	DEBUG: fmt.Sprintf(colorSeq, 34), // blue
	WARN:  fmt.Sprintf(colorSeq, 33), // yellow
	ERROR: fmt.Sprintf(colorSeq, 35), // pink
	FATAL: fmt.Sprintf(colorSeq, 31), // red
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
