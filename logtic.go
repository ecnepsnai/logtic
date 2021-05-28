// Package logtic is a yet another logging library for golang projects.
//
// The goal of logtic is to be as transparent and easy to use as possible, allowing applications and libraries to
// seamlessly log to a single file. Logtic can be used in libraries and won't cause any problems if the parent
// application isn't using logtic.
//
// Logtic supports multiple sources, which annotate the outputted log lines. It also supports defining a minimum
// desired log level, which can be changed at any time. Events printed to the terminal output support color-coded
// severities.
//
// Events can be printed as formatted strings, like with `fmt.Printf`, or can be parameterized events which can be easily
// parsed by log analysis tools such as Splunk.
//
// By default, logtic will only print to stdout and stderr, but when configured it can also write to a log file. Log files
// include the date-time for each event in RFC-3339 format. Log files can be rotated using the `logtic.Rotate()` method.
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
	Color:    true,
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

// Reset will logtic to an unconfigured state, closing any open log files.
func Reset() {
	Close()
	Log = &Settings{
		FilePath: os.DevNull,
		Level:    LevelError,
		FileMode: 0644,
		lock:     sync.Mutex{},
		Color:    true,
	}
}

// Connect will prepare a new logtic source with the given name. Sources can be written even if there is no open logtic
// log session.
func Connect(sourceName string) *Source {
	return &Source{
		Name:  sourceName,
		Level: -1,
	}
}

// Close will the log file.
func Close() {
	if Log.file != nil {
		Log.file.Close()
	}
}
