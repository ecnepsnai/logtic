package logtic

import (
	"os"
	"time"
)

// File describes a log file
type File struct {
	file *os.File
}

// Close close the log file
func (f *File) Close() {
	if f != nil && f.file != nil {
		f.file.Close()
	}
	deleteInstance()
}

func (f File) write(message string) {
	f.file.WriteString(time.Now().Format(time.RFC3339) + " " + message + "\n")
}
