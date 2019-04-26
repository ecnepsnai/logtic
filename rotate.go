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
func (f *File) Rotate() error {
	f.lock.Lock()
	defer f.lock.Unlock()

	extensionComponents := strings.Split(f.logPath, ".")
	extension := extensionComponents[len(extensionComponents)-1]

	pathComponents := strings.Split(f.logPath, string(os.PathSeparator))
	fileName := pathComponents[len(pathComponents)-1]
	date := time.Now().Format("2006-01-02")
	newName := strings.Replace(fileName, extension, date+"."+extension, -1)
	newPath := strings.Replace(f.logPath, fileName, newName, -1)

	f.file.Close()
	if err := os.Rename(f.logPath, newPath); err != nil {
		return err
	}
	newFile, err := os.OpenFile(f.logPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	f.file = newFile
	return nil
}
