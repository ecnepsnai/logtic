package logtic

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Rotate will rotate the logfile. The current logfile will be renamed and suffixed with the current date in a
// YYYY-MM-DD format. A new log file will be opened with the original file path and used for all subsequent
// writes. Writes will be blocked while the rotation is in progress. If a file matching the name of what would be used
// for the rotated file, a numerical suffix is added to the end of the name.
//
// If an error is returned during rotation it is highly recommended that you either panic or call logtic.Reset()
// as logtic may be in an undefined state and log calls may cause panics.
func Rotate() error {
	if Log.file == nil {
		return nil
	}

	Log.lock.Lock()
	defer Log.lock.Unlock()

	extensionComponents := strings.Split(Log.FilePath, ".")
	extension := extensionComponents[len(extensionComponents)-1]

	pathComponents := strings.Split(Log.FilePath, string(os.PathSeparator))
	fileName := pathComponents[len(pathComponents)-1]
	date := time.Now().Format("2006-01-02")
	newName := strings.Replace(fileName, extension, extension+"."+date, -1)
	newPath := strings.Replace(Log.FilePath, fileName, newName, -1)

	if fileExists(newPath) {
		i := 1
		for fileExists(fmt.Sprintf("%s-%d", newPath, i)) {
			i++
		}
		newPath = fmt.Sprintf("%s-%d", newPath, i)
	}

	Log.file.Close()
	Log.file = nil

	if err := os.Rename(Log.FilePath, newPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error renaming existing log file: %s", err.Error())
		return err
	}
	if err := Open(); err != nil {
		fmt.Fprintf(os.Stderr, "Error opening new log file '%s': %s", Log.FilePath, err.Error())
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
