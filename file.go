package logtic

import (
	"os"
	"sync"
	"time"
)

// File describes a log file
type File struct {
	file    *os.File
	logPath string
	lock    sync.Mutex
}

// Close close the log file
func (f *File) Close() {
	if f != nil && f.file != nil {
		f.file.Close()
	}
	deleteInstance()
}

func (f *File) write(message string) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.file.WriteString(time.Now().Format(time.RFC3339) + " " + message + "\n")
}
