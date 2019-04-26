// Package logtic is a (another) logging library for golang projects
//
// Logtic is meant for large applications that contain multiple libraries
// that all need to write to a single log file.
// Logtic is transparent in that it can be included in your libraries and attach
// to any log file if the parent application is using logtic, otherwise it
// just does nothing.
package logtic

import (
	"os"
	"sync"
	"unsafe"
)

// New create a new logtic log file and source. This should only be called by
// the running application once, at launch.
func New(path string, level int, sourceName string) (*File, *Source, error) {
	settings, err := discoverDescriptorForProcess()
	if err != nil {
		return nil, nil, err
	}
	var logFile *os.File
	if settings != nil {
		pointer := uintptr(settings.FilePointer)
		logFile = (*os.File)(unsafe.Pointer(pointer))
	} else {
		f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			return nil, nil, err
		}
		logFile = f

		pointer := unsafe.Pointer(f)

		s := instance{
			FilePointer: uint64(uintptr(pointer)),
			Path:        path,
			Level:       level,
		}
		s.save()
	}

	file := File{
		file:    logFile,
		logPath: path,
		lock:    sync.Mutex{},
	}
	source := Source{
		file:      file,
		className: sourceName,
		Level:     level,
	}

	return &file, &source, nil
}

// Connect connect to an existing logtic log file for this process and inherit its settings
// if no logtic session is running, do nothing
func Connect(sourceName string) *Source {
	settings, err := discoverDescriptorForProcess()
	dummy := dummySource()
	if err != nil {
		return &dummy
	}
	var logFile *os.File
	if settings != nil {
		pointer := uintptr(settings.FilePointer)
		logFile = (*os.File)(unsafe.Pointer(pointer))
	} else {
		return &dummy
	}

	file := File{
		file: logFile,
	}
	source := Source{
		file:      file,
		className: sourceName,
		Level:     settings.Level,
	}

	return &source
}
