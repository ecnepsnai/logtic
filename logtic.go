// Package logtic is a (another) logging library for golang projects
//
// Logtic is meant for large applications that contain multiple libraries
// that all need to write to a single log file.
// Logtic is transparent in that it can be included in your libraries and attach
// to any log file if the parent application is using logtic, otherwise it
// just does nothing.
// The overall goal of logtic is that "it just works", meaning there should be little
// effort required to get it working the correct way.
package logtic

import (
	"os"
	"sync"
)

// Log the global log settings for this application
var Log = &Settings{
	lock: sync.Mutex{},
}

// Open will open the file specified by the FilePath for writing. It will create the file if it does not
// exist, and append to existing files.
func Open() error {
	if Log.file != nil {
		return nil
	}

	f, err := os.OpenFile(Log.FilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
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
		lock: sync.Mutex{},
	}
}

// Connect connect to an existing logtic log file for this process and inherit its settings
// if no logtic session is running, do nothing
func Connect(sourceName string) *Source {
	dummy := dummySource()
	if Log.file == nil {
		return &dummy
	}
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
