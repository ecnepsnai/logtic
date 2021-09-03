package logtic_test

import (
	"path"
	"sync"
	"testing"

	"github.com/ecnepsnai/logtic"
)

func TestWrite(t *testing.T) {
	logtic.Log.Reset()

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

func TestDummy(t *testing.T) {
	logtic.Log.Reset()

	var wg sync.WaitGroup

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
	logtic.Log.Reset()

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
	logtic.Log.Reset()

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
