package logtic

import (
	"os"
	"sync"
)

// Settings describes the log settings for this application
type Settings struct {
	// FilePath the path to the log file.
	FilePath string
	// Level the log level. Use the predefined logtic.Level(X) constants.
	Level int

	file *os.File
	lock sync.Mutex
}
