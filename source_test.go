package logtic_test

import (
	"io"
	"os"
	"path"
	"regexp"
	"testing"

	"github.com/ecnepsnai/logtic"
)

// Test that the expected lines are printed to a log file
func TestSources(t *testing.T) {
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

	source := logtic.Connect("test")

	source.Write(logtic.LevelDebug, "this is a %s message", "debug")
	debugPattern := regexp.MustCompile(`[0-9\-:T]+ \[DEBUG\]\[test\] this is a debug message`)
	source.Write(logtic.LevelInfo, "this is an %s message", "info")
	infoPattern := regexp.MustCompile(`[0-9\-:T]+ \[INFO\]\[test\] this is an info message`)
	source.Write(logtic.LevelWarn, "this is a %s message", "warning")
	warnPattern := regexp.MustCompile(`[0-9\-:T]+ \[WARN\]\[test\] this is a warning message`)
	source.Write(logtic.LevelError, "this is an %s message", "error")
	errorPattern := regexp.MustCompile(`[0-9\-:T]+ \[ERROR\]\[test\] this is an error message`)

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

	if !debugPattern.Match(logFileData) {
		t.Errorf("Log file does not contain expected log line for Debug message")
	}
	if !infoPattern.Match(logFileData) {
		t.Errorf("Log file does not contain expected log line for Info message")
	}
	if !warnPattern.Match(logFileData) {
		t.Errorf("Log file does not contain expected log line for Warn message")
	}
	if !errorPattern.Match(logFileData) {
		t.Errorf("Log file does not contain expected log line for Error message")
	}
}

func TestPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

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

	source := logtic.Connect("test")

	source.Panic("Ahh!")
}

func TestSourceLevel(t *testing.T) {
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

	source1 := logtic.Connect("source1")
	source2 := logtic.Connect("source2")
	source2.Level = logtic.LevelWarn

	source1.Info("info message")
	source2.Info("info message")

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

	source1Pattern := regexp.MustCompile(`[0-9\-:T]+ \[INFO\]\[source1\] info message`)
	source2Pattern := regexp.MustCompile(`[0-9\-:T]+ \[INFO\]\[source2\] info message`)
	if !source1Pattern.Match(logFileData) {
		t.Errorf("Log file does not contain expected log lines")
	}
	if source2Pattern.Match(logFileData) {
		t.Errorf("Log file contains log line that should not exist")
	}
}

func TestSourceWriteUnknownLevel(t *testing.T) {
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

	source := logtic.Connect("test")

	source.Write(9001, "What level is this?")
	source.PWrite(9001, "Crazy!", nil)
}
