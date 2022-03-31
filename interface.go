package logh

// LogH will accept stdlib logger and a custom logger.
// There's no standard interface, this is the closest we get, unfortunately.
type LogH interface {
	Print(...interface{})
	Printf(string, ...interface{})

	Fatal(...interface{})
	Fatalf(string, ...interface{})

	Panic(...interface{})
	Panicf(string, ...interface{})
}
