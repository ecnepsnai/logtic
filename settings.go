package logtic

import (
	"os"
	"sync"
)

// Settings describes the log settings for this application
type Settings struct {
	// The path to the log file.
	FilePath string
	// The minimum level of events captured in the log file and printed to console. Inclusive.
	Level int
	// The file mode (permissions) used for the log file and rotated log files.
	FileMode os.FileMode
	// Should logtic use color for events printed to stdout/stderr
	Color bool

	file *os.File
	lock sync.Mutex
}
