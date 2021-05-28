package logtic_test

import (
	"io"
	"os"
	"path"
	"regexp"
	"testing"

	"github.com/ecnepsnai/logtic"
)

func TestSourceParameters(t *testing.T) {
	logtic.Reset()

	dir, err := os.MkdirTemp("", "logtic")
	if err != nil {
		panic(err)
	}

	logtic.Log.FilePath = path.Join(dir, "logtic.log")
	logtic.Log.Level = logtic.LevelDebug

	if err := logtic.Open(); err != nil {
		t.Fatalf("Error opening log file: %s", err.Error())
	}

	parameters := map[string]interface{}{
		"string": "hello, world!",
		"int":    123,
		"float":  3.14,
		"bool":   true,
		"bytes":  []byte("Hello, world!"),
		"slice":  []int{1, 2, 3},
	}
	source := logtic.Connect("test")
	source.PWrite(logtic.LevelDebug, "Event", parameters)
	source.PWrite(logtic.LevelInfo, "Event", parameters)
	source.PWrite(logtic.LevelWarn, "Event", parameters)
	source.PWrite(logtic.LevelError, "Event", parameters)
	pattern := regexp.MustCompile(`[0-9\-:T]+ \[(DEBUG|INFO|WARN|ERROR)\]\[test\] Event: bool='true' bytes=48656c6c6f2c20776f726c6421 float=3.140000 int=123 slice='\[1 2 3\]' string='hello, world!'`)

	logtic.Close()

	f, err := os.OpenFile(path.Join(dir, "logtic.log"), os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	logFileData, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	if !pattern.Match(logFileData) {
		t.Errorf("Log file does not contain expected log line for Debug message")
	}
}
