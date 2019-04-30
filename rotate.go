package logtic

import (
	"os"
	"strings"
	"time"
)

// Rotate will rotate the logfile. The current logfile will be renamed and
// suffixed with the current date in a YYYY-MM-DD format.
// A new log file will be opened with the same original file and used for all
// subsequent writes. Writes will be blocked while the rotation is in progress.
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
	newName := strings.Replace(fileName, extension, date+"."+extension, -1)
	newPath := strings.Replace(Log.FilePath, fileName, newName, -1)

	Log.file.Close()
	if err := os.Rename(Log.FilePath, newPath); err != nil {
		return err
	}

	newFile, err := os.OpenFile(Log.FilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	Log.file = newFile
	return nil
}
