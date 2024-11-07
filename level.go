package logtic

type LogLevel int

const (
	// LevelDebug debug messages for troubleshooting application behaviour
	LevelDebug = LogLevel(3)
	// LevelInfo informational messages for normal operation of the application
	LevelInfo = LogLevel(2)
	// LevelWarn warning messages for potential issues
	LevelWarn = LogLevel(1)
	// LevelError error messages for problems
	LevelError = LogLevel(0)
)
