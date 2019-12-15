package logtic

import (
	"fmt"
	"io"
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
	newName := strings.Replace(fileName, extension, extension+"."+date, -1)
	newPath := strings.Replace(Log.FilePath, fileName, newName, -1)

	// Rewind the file pointer
	_, err := Log.file.Seek(0, io.SeekStart)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error rewinding log file '%s': %s\n", Log.FilePath, err.Error())
		return err
	}

	// Open the new log file
	newFile, err := os.OpenFile(newPath, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening rotated log file '%s': %s\n", newPath, err.Error())
		return err
	}

	// Copy the contents and close the file
	length, err := io.Copy(newFile, Log.file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error copying current log file '%s' contents to rotated file '%s': %s\n", Log.FilePath, newPath, err.Error())
		return err
	}
	newFile.Close()

	// Truncate the file
	Log.file.Seek(length, io.SeekEnd)
	if err := Log.file.Truncate(0); err != nil {
		fmt.Fprintf(os.Stderr, "Error truncating current log file '%s': %s\n", Log.FilePath, err.Error())
		return err
	}

	return nil
}
