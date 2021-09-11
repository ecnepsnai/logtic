package logtic

import (
	"os"
	"sync"
	"time"
)

// Logger describes a logging instance
type Logger struct {
	// The path to the log file.
	FilePath string
	// The minimum level of events captured in the log file and printed to console. Inclusive.
	Level int
	// The file mode (permissions) used for the log file and rotated log files.
	FileMode os.FileMode
	// Options various options for this logger
	Options LoggerOptions

	file *os.File
	lock sync.Mutex
}

// LoggerOptions describe logger options
type LoggerOptions struct {
	// Should logtic not use color when printing events to stdout/stderr. Enabled by default.
	Color bool
	// Should logtic escape control characters automatically. For example, replaces actual newlines with a literal \n.
	// Enabled by default.
	EscapeCharacters bool
	// Should logtic automatically gzip log files after rotation. Disabled by default.
	GZipRotatedLogs bool
}

func defaultLoggerOption() LoggerOptions {
	return LoggerOptions{
		Color:            true,
		EscapeCharacters: true,
		GZipRotatedLogs:  false,
	}
}

// Log is the default logging instance.
var Log = New()

// New will create a new logging instance with all default options. You should only use new if you want a separate
// logging instance from the default instance, which is automatically created for you.
func New() *Logger {
	return &Logger{
		FilePath: os.DevNull,
		Level:    LevelError,
		FileMode: 0644,
		Options:  defaultLoggerOption(),
	}
}

// Open will open the file specified by FilePath on this logging instance. The file will be created if it does
// not already exist, otherwise it will be appended to.
func (l *Logger) Open() error {
	if l.file != nil {
		return nil
	}

	f, err := os.OpenFile(l.FilePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, l.FileMode)
	if err != nil {
		return err
	}
	l.file = f

	return nil
}

// Reset will reset this logging instance to its original state. Open files will be closed.
func (l *Logger) Reset() {
	l.Close()
	l.FilePath = os.DevNull
	l.Level = LevelError
	l.FileMode = 0644
	l.lock = sync.Mutex{}
	l.Options = defaultLoggerOption()
	l.file = nil
}

// Connect will prepare a new logtic source with the given name for this logging instance. Sources can be written
// even if there is no open logtic log instance.
func (l *Logger) Connect(sourceName string) *Source {
	return &Source{
		Name:     sourceName,
		Level:    -1,
		instance: l,
	}
}

// Close will flush and close this logging instance.
func (l *Logger) Close() {
	if l.file != nil {
		l.file.Sync()
		l.file.Close()
		l.file = nil
	}
}

func (l *Logger) write(message string) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.file.WriteString(time.Now().Format(time.RFC3339) + " " + message + "\n")
}
