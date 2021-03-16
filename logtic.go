// Package logtic is a (another) logging library for golang projects.
//
// The goal of logtic is to be as transparent and easy to use as possible, allowing applications and libraries to
// seamlessly log to a single file. Logtic can be used in libraries and won't cause any problems if the parent
// application isn't using logtic.
//
// Logtic supports multiple sources, which annotate the outputted log lines. It also supports defining a minimum
// desired log level, which can be changed at any time.
//
// By default, logtic will only print to stdout and stderr, but when configured it can also write to a log file.
// Logtic can also rotate these log files out by invoking the logtic.Rotate() method. Log files include the date-time
// for each line in RFC-3339 format.
package logtic

import (
	"os"
	"sync"
)

// Log the global log settings for this application
var Log = &Settings{
	FilePath: os.DevNull,
	Level:    LevelError,
	FileMode: 0644,
}

// Open will open the file specified by the FilePath for writing. It will create the file if it does not
// exist, and append to existing files.
func Open() error {
	if Log.file != nil {
		return nil
	}

	f, err := os.OpenFile(Log.FilePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, Log.FileMode)
	if err != nil {
		return err
	}
	Log.file = f

	return nil
}

// Reset reset logtic to an unconfigured state, closing any open log files
func Reset() {
	Close()
	Log = &Settings{
		FilePath: os.DevNull,
		Level:    LevelError,
		FileMode: 0644,
		lock:     sync.Mutex{},
	}
}

// Connect connect to an existing logtic log file for this process and inherit its settings
// if no logtic session is running, do nothing
func Connect(sourceName string) *Source {
	return &Source{
		className: sourceName,
	}
}

// Close close the log file
func Close() {
	if Log.file != nil {
		Log.file.Close()
	}
}
