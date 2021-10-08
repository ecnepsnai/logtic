package logtic_test

import (
	"io"
	"os"
	"path"
	"sync"
	"testing"

	"github.com/ecnepsnai/logtic"
)

var verbose = false

func TestMain(m *testing.M) {
	for _, arg := range os.Args {
		if arg == "-test.v=true" {
			verbose = true
			break
		}
	}

	os.Exit(m.Run())
}

func SetStdOut(log *logtic.Logger) {
	if verbose {
		log.Stdout = os.Stdout
		log.Stderr = os.Stderr
	} else {
		log.Stdout = io.Discard
		log.Stderr = io.Discard
	}
}

func Setup() {
	logtic.Log.Reset()
	SetStdOut(logtic.Log)
}

func TestWrite(t *testing.T) {
	Setup()

	dir := t.TempDir()

	logtic.Log.FilePath = path.Join(dir, "logtic.log")
	logtic.Log.Level = logtic.LevelDebug

	if err := logtic.Log.Open(); err != nil {
		t.Fatalf("Error opening log file: %s", err.Error())
	}

	s := logtic.Log.Connect("Test")

	var wg sync.WaitGroup

	s.Debug("Start test")

	wg.Add(3)
	go func() {
		defer wg.Done()
		source := logtic.Log.Connect("goroutine1")
		i := 0
		for i < 5 {
			i++
			source.Debug("Count %d", i)
		}
	}()
	go func() {
		defer wg.Done()
		source := logtic.Log.Connect("goroutine2")
		i := 0
		for i < 5 {
			i++
			source.Debug("Count %d", i)
		}
	}()
	go func() {
		defer wg.Done()
		source := logtic.Log.Connect("goroutine3")
		i := 0
		for i < 5 {
			i++
			source.Debug("Count %d", i)
		}
	}()

	wg.Wait()
}

func TestEarlyConnect(t *testing.T) {
	Setup()

	s := logtic.Log.Connect("Test")

	dir := t.TempDir()

	logtic.Log.FilePath = path.Join(dir, "logtic.log")
	logtic.Log.Level = logtic.LevelDebug

	if err := logtic.Log.Open(); err != nil {
		t.Fatalf("Error opening log file: %s", err.Error())
	}

	var wg sync.WaitGroup

	s.Debug("Start test")

	wg.Add(3)
	go func() {
		defer wg.Done()
		source := logtic.Log.Connect("goroutine1")
		i := 0
		for i < 5 {
			i++
			source.Debug("Count %d", i)
		}
	}()
	go func() {
		defer wg.Done()
		source := logtic.Log.Connect("goroutine2")
		i := 0
		for i < 5 {
			i++
			source.Debug("Count %d", i)
		}
	}()
	go func() {
		defer wg.Done()
		source := logtic.Log.Connect("goroutine3")
		i := 0
		for i < 5 {
			i++
			source.Debug("Count %d", i)
		}
	}()

	wg.Wait()
}

func TestOpenTwice(t *testing.T) {
	Setup()

	s := logtic.Log.Connect("Test")

	dir := t.TempDir()

	logtic.Log.FilePath = path.Join(dir, "logtic.log")
	logtic.Log.Level = logtic.LevelDebug

	if err := logtic.Log.Open(); err != nil {
		t.Fatalf("Error opening log file: %s", err.Error())
	}

	s.Debug("Testing 123")

	if err := logtic.Log.Open(); err != nil {
		t.Fatalf("Unexpected error opening already open log file: %s", err.Error())
	}

	s2 := logtic.Log.Connect("Test2")
	s2.Debug("Testing 123")
}
