package logtic_test

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"testing"

	"github.com/ecnepsnai/logtic"
)

// Test that the expected lines are printed to a log file
func TestSources(t *testing.T) {
	logtic.Log.Reset()

	dir := t.TempDir()

	logtic.Log.FilePath = path.Join(dir, "logtic.log")
	logtic.Log.Level = logtic.LevelDebug

	if err := logtic.Log.Open(); err != nil {
		t.Fatalf("Error opening log file: %s", err.Error())
	}

	source := logtic.Log.Connect("test")

	source.Write(logtic.LevelDebug, "this is a %s message", "debug")
	debugPattern := regexp.MustCompile(`[0-9\-:TZ]+ \[DEBUG\]\[test\] this is a debug message`)
	source.Write(logtic.LevelInfo, "this is an %s message", "info")
	infoPattern := regexp.MustCompile(`[0-9\-:TZ]+ \[INFO\]\[test\] this is an info message`)
	source.Write(logtic.LevelWarn, "this is a %s message", "warning")
	warnPattern := regexp.MustCompile(`[0-9\-:TZ]+ \[WARN\]\[test\] this is a warning message`)
	source.Write(logtic.LevelError, "this is an %s message", "error")
	errorPattern := regexp.MustCompile(`[0-9\-:TZ]+ \[ERROR\]\[test\] this is an error message`)

	logtic.Log.Close()

	logFileData, err := os.ReadFile(path.Join(dir, "logtic.log"))
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
	if t.Failed() {
		fmt.Printf("Log file data:\n%s\n", logFileData)
	}
}

func TestPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	logtic.Log.Reset()

	dir := t.TempDir()

	logtic.Log.FilePath = path.Join(dir, "logtic.log")
	logtic.Log.Level = logtic.LevelDebug

	if err := logtic.Log.Open(); err != nil {
		t.Fatalf("Error opening log file: %s", err.Error())
	}

	source := logtic.Log.Connect("test")

	source.Panic("Ahh!")
}

func TestSourceLevel(t *testing.T) {
	logtic.Log.Reset()

	dir := t.TempDir()

	logtic.Log.FilePath = path.Join(dir, "logtic.log")
	logtic.Log.Level = logtic.LevelDebug

	if err := logtic.Log.Open(); err != nil {
		t.Fatalf("Error opening log file: %s", err.Error())
	}

	source1 := logtic.Log.Connect("source1")
	source2 := logtic.Log.Connect("source2")
	source2.Level = logtic.LevelWarn

	source1.Info("info message")
	source2.Info("info message")

	logtic.Log.Close()

	logFileData, err := os.ReadFile(path.Join(dir, "logtic.log"))
	if err != nil {
		panic(err)
	}

	source1Pattern := regexp.MustCompile(`[0-9\-:TZ]+ \[INFO\]\[source1\] info message`)
	source2Pattern := regexp.MustCompile(`[0-9\-:TZ]+ \[INFO\]\[source2\] info message`)
	if !source1Pattern.Match(logFileData) {
		t.Errorf("Log file does not contain expected log lines")
	}
	if source2Pattern.Match(logFileData) {
		t.Errorf("Log file contains log line that should not exist")
	}
}

func TestSourceWriteUnknownLevel(t *testing.T) {
	logtic.Log.Reset()

	dir := t.TempDir()

	logtic.Log.FilePath = path.Join(dir, "logtic.log")
	logtic.Log.Level = logtic.LevelDebug

	if err := logtic.Log.Open(); err != nil {
		t.Fatalf("Error opening log file: %s", err.Error())
	}

	source := logtic.Log.Connect("test")

	source.Write(9001, "What level is this?")
	source.PWrite(9001, "Crazy!", nil)
}

func TestMultipleLogs(t *testing.T) {
	logtic.Log.Reset()

	dir := t.TempDir()

	aLogPath := path.Join(dir, "a.log")
	bLogPath := path.Join(dir, "b.log")

	logtic.Log.FilePath = aLogPath
	logtic.Log.Level = logtic.LevelDebug

	if err := logtic.Log.Open(); err != nil {
		t.Fatalf("Error opening a log file: %s", err.Error())
	}

	aSource := logtic.Log.Connect("a-source")

	bLog := logtic.New()
	bLog.FilePath = bLogPath
	bLog.Level = logtic.LevelDebug

	if err := bLog.Open(); err != nil {
		t.Fatalf("Error opening b log file: %s", err.Error())
	}

	bSource := bLog.Connect("b-source")

	aSource.Info("Info")
	bSource.Warn("Warn")

	logtic.Log.Close()
	bLog.Close()

	aLogData, err := os.ReadFile(aLogPath)
	if err != nil {
		t.Fatalf("Error reading log file data: %s", err.Error())
	}
	bLogData, err := os.ReadFile(bLogPath)
	if err != nil {
		t.Fatalf("Error reading log file data: %s", err.Error())
	}

	infoPattern := regexp.MustCompile(`\[INFO\]\[[ab]-source\] Info`)
	warnPattern := regexp.MustCompile(`\[WARN\]\[[ab]-source\] Warn`)
	if !infoPattern.Match(aLogData) {
		t.Errorf("Incorrect log data sent to source: info not found in A log")
	}
	if warnPattern.Match(aLogData) {
		t.Errorf("Incorrect log data sent to source: warn found in A log")
	}
	if infoPattern.Match(bLogData) {
		t.Errorf("Incorrect log data sent to source: info found in B log")
	}
	if !warnPattern.Match(bLogData) {
		t.Errorf("Incorrect log data sent to source: warn not found in B log")
	}

	if t.Failed() {
		t.Logf("a file: %s", aLogData)
		t.Logf("b file: %s", bLogData)
	}
}

func TestSourceEscapeCharacters(t *testing.T) {
	logtic.Log.Reset()

	dir := t.TempDir()

	logtic.Log.FilePath = path.Join(dir, "logtic.log")
	logtic.Log.Level = logtic.LevelDebug

	if err := logtic.Log.Open(); err != nil {
		t.Fatalf("Error opening log file: %s", err.Error())
	}

	source := logtic.Log.Connect("test")

	source.Write(logtic.LevelDebug, "Hello\n%s", "world")
	debugPattern := regexp.MustCompile(`[0-9\-:TZ]+ \[DEBUG\]\[test\] Hello\\nworld`)

	logtic.Log.Options.EscapeCharacters = false

	source.Write(logtic.LevelInfo, "Hello\n%s", "world")
	infoPattern := regexp.MustCompile(`[0-9\-:TZ]+ \[INFO\]\[test\] Hello\nworld`)

	logtic.Log.Close()

	logFileData, err := os.ReadFile(path.Join(dir, "logtic.log"))
	if err != nil {
		panic(err)
	}

	if !debugPattern.Match(logFileData) {
		t.Errorf("Log file does not contain expected log line for Debug message")
	}
	if !infoPattern.Match(logFileData) {
		t.Errorf("Log file does not contain expected log line for Info message")
	}
	if t.Failed() {
		fmt.Printf("Log file data:\n%s\n", logFileData)
	}
}
