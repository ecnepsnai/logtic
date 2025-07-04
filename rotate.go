package logtic

import (
	"fmt"
	"os"
	"time"
)

// Rotate will rotate the log file of this logging instance. The current log file will be renamed and suffixed
// with the current date in a YYYY-MM-DD format. A new log file will be opened with the original file path and used for
// all subsequent writes. Writes will be blocked while the rotation is in progress. If a file matching the name of what
// would be used for the rotated file, a dash and numerical suffix is added to the end of the name.
//
// If an error is returned during rotation it is highly recommended that you either panic or call logger.Reset()
// as logtic may be in an undefined state and log calls may cause panics.
//
// If no log file has been opened on this logger, calls to Rotate do nothing.
func (l *Logger) Rotate() error {
	if l.file == nil {
		return nil
	}

	l.lock.Lock()
	defer l.lock.Unlock()

	date := time.Now().Format("2006-01-02")
	newPath := l.FilePath + "." + date

	if fileExists(newPath) {
		i := 1
		for fileExists(fmt.Sprintf("%s-%d", newPath, i)) {
			i++
		}
		newPath = fmt.Sprintf("%s-%d", newPath, i)
	}

	if err := l.file.Sync(); err != nil {
		fmt.Fprintf(os.Stderr, "Error syncing changes to existing log file: %s", err.Error())
		return err
	}
	if err := l.file.Close(); err != nil {
		fmt.Fprintf(os.Stderr, "Error closing existing log file: %s", err.Error())
		return err
	}
	l.file = nil

	if err := os.Rename(l.FilePath, newPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error renaming existing log file: %s", err.Error())
		return err
	}
	if err := Log.Open(); err != nil {
		fmt.Fprintf(os.Stderr, "Error opening new log file '%s': %s", l.FilePath, err.Error())
		return err
	}

	return nil
}

func fileExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}
