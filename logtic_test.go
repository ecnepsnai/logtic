package logtic_test

import (
	"io/ioutil"
	"path"
	"sync"
	"testing"

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
