package logtic_test

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"regexp"
	"testing"

	"github.com/ecnepsnai/logtic"
)

// Test that the expected lines are printed to a log file
func TestSources(t *testing.T) {
	Setup()

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

func TestPanicEvent(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	Setup()

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
	Setup()

	dir := t.TempDir()

	logtic.Log.FilePath = path.Join(dir, "logtic.log")
	logtic.Log.Level = logtic.LevelDebug

	if err := logtic.Log.Open(); err != nil {
		t.Fatalf("Error opening log file: %s", err.Error())
	}

	source1 := logtic.Log.Connect("source1")
	source2 := logtic.Log.Connect("source2")
	source2.OverrideLevel(logtic.LevelWarn)

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
	Setup()

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
	Setup()

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
	SetStdOut(bLog)
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
	Setup()

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

// Test that nothing is printed to the console when the logger is not opened
func TestLoggerNotOpened(t *testing.T) {
	Setup()

	b := &bytes.Buffer{}
	logtic.Log.Stdout = b
	logtic.Log.Stderr = b

	source := logtic.Log.Connect("example")
	source.Debug("Debug")
	source.Info("Info")
	source.Warn("Warn")
	source.Error("Error")

	if b.Len() != 0 {
		t.Errorf("Data written to stdout/stderr that wasn't expected")
	}

	// Panic and Fatal messages are always printed to stderr
	defer func() {
		if r := recover(); r != nil {
			Setup()

			if b.Len() == 0 {
				t.Errorf("Panic message not printed to stderr when expected")
			}
		}
	}()
	source.Panic("Panic!")
}

// Test that the instance does not panic (even if recovered) if a write event is called on a nil source or nil instance
func TestNilLogInstance(t *testing.T) {
	origStderr := os.Stderr
	tempStderr := path.Join(t.TempDir(), "stderr")
	f, err := os.OpenFile(tempStderr, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	os.Stderr = f

	var log *logtic.Source
	log.Error("Error")

	f.Close()

	data, err := os.ReadFile(tempStderr)
	if err != nil {
		panic(err)
	}
	if bytes.Contains(data, []byte("logtic: recovered from panic writing event. stack to follow")) {
		t.Errorf("logtic instance did panic")
		t.Logf("%s", data)
	}

	os.Stderr = origStderr
}
