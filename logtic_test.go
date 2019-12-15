package logtic_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sync"
	"testing"
	"time"

	"github.com/ecnepsnai/logtic"
)

func TestWrite(t *testing.T) {
	logtic.Reset()

	dir, err := ioutil.TempDir("", "logtic")
	if err != nil {
		panic(err)
	}

	logtic.Log.FilePath = path.Join(dir, "logtic.log")
	logtic.Log.Level = logtic.LevelDebug

	if err := logtic.Open(); err != nil {
		t.Fatalf("Error opening log file: %s", err.Error())
	}

	s := logtic.Connect("Test")

	var wg sync.WaitGroup

	s.Debug("Start test")

	wg.Add(3)
	go func() {
		defer wg.Done()
		source := logtic.Connect("goroutine1")
		i := 0
		for i < 100 {
			i++
			source.Debug("Count %d", i)
			source.Info("Count %d", i)
			source.Warn("Count %d", i)
			source.Error("Count %d", i)
		}
	}()
	go func() {
		defer wg.Done()
		source := logtic.Connect("goroutine2")
		i := 0
		for i < 100 {
			i++
			source.Debug("Count %d", i)
			source.Info("Count %d", i)
			source.Warn("Count %d", i)
			source.Error("Count %d", i)
		}
	}()
	go func() {
		defer wg.Done()
		source := logtic.Connect("goroutine3")
		i := 0
		for i < 100 {
			i++
			source.Debug("Count %d", i)
			source.Info("Count %d", i)
			source.Warn("Count %d", i)
			source.Error("Count %d", i)
		}
	}()

	wg.Wait()
}

func TestDummy(t *testing.T) {
	logtic.Reset()

	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		defer wg.Done()
		source := logtic.Connect("goroutine1")
		i := 0
		for i < 100 {
			i++
			source.Debug("Count %d", i)
			source.Info("Count %d", i)
			source.Warn("Count %d", i)
			source.Error("Count %d", i)
		}
	}()
	go func() {
		defer wg.Done()
		source := logtic.Connect("goroutine2")
		i := 0
		for i < 100 {
			i++
			source.Debug("Count %d", i)
			source.Info("Count %d", i)
			source.Warn("Count %d", i)
			source.Error("Count %d", i)
		}
	}()
	go func() {
		defer wg.Done()
		source := logtic.Connect("goroutine3")
		i := 0
		for i < 100 {
			i++
			source.Debug("Count %d", i)
			source.Info("Count %d", i)
			source.Warn("Count %d", i)
			source.Error("Count %d", i)
		}
	}()

	wg.Wait()
}

func TestRotate(t *testing.T) {
	logtic.Reset()

	dir, err := ioutil.TempDir("", "logtic")
	if err != nil {
		panic(err)
	}

	logtic.Log.FilePath = path.Join(dir, "app.log")
	logtic.Log.Level = logtic.LevelDebug

	if err := logtic.Open(); err != nil {
		t.Fatalf("Error opening log file: %s", err.Error())
	}

	s := logtic.Connect("Test")
	i := 0
	for i < 5 {
		i++
		s.Debug("Count %d", i)
		s.Info("Count %d", i)
		s.Warn("Count %d", i)
		s.Error("Count %d", i)
	}

	if err := logtic.Rotate(); err != nil {
		panic(err)
	}

	s.Info("Rotated log")

	date := time.Now().Format("2006-01-02")
	expectedPath := path.Join(dir, "app.log."+date)
	if _, err := os.Stat(expectedPath); err != nil {
		t.Errorf("Expected rotated log file not found: '%s'", expectedPath)
	}

	currentPath := path.Join(dir, "app.log")
	if _, err := os.Stat(currentPath); err != nil {
		t.Errorf("Expected new log file not found: '%s'", currentPath)
	}

	checkFileSize := func(path string) {
		info, err := os.Stat(path)
		if err != nil {
			t.Errorf("Error stating rotated log file: %s", err.Error())
		}

		fmt.Printf("%s -> %d\n", path, info.Size())

		if info.Size() == 0 {
			t.Errorf("Rotated log file is empty")
		}
	}

	checkFileSize(expectedPath)
	checkFileSize(currentPath)
}

func TestEarlyConnect(t *testing.T) {
	logtic.Reset()

	s := logtic.Connect("Test")

	dir, err := ioutil.TempDir("", "logtic")
	if err != nil {
		panic(err)
	}

	logtic.Log.FilePath = path.Join(dir, "logtic.log")
	logtic.Log.Level = logtic.LevelDebug

	if err := logtic.Open(); err != nil {
		t.Fatalf("Error opening log file: %s", err.Error())
	}

	var wg sync.WaitGroup

	s.Debug("Start test")

	wg.Add(3)
	go func() {
		defer wg.Done()
		source := logtic.Connect("goroutine1")
		i := 0
		for i < 100 {
			i++
			source.Debug("Count %d", i)
			source.Info("Count %d", i)
			source.Warn("Count %d", i)
			source.Error("Count %d", i)
		}
	}()
	go func() {
		defer wg.Done()
		source := logtic.Connect("goroutine2")
		i := 0
		for i < 100 {
			i++
			source.Debug("Count %d", i)
			source.Info("Count %d", i)
			source.Warn("Count %d", i)
			source.Error("Count %d", i)
		}
	}()
	go func() {
		defer wg.Done()
		source := logtic.Connect("goroutine3")
		i := 0
		for i < 100 {
			i++
			source.Debug("Count %d", i)
			source.Info("Count %d", i)
			source.Warn("Count %d", i)
			source.Error("Count %d", i)
		}
	}()

	wg.Wait()
}
