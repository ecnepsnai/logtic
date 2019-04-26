package logtic

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sync"
	"testing"
	"time"
)

func TestWrite(t *testing.T) {
	dir, err := ioutil.TempDir("", "logtic")
	if err != nil {
		fmt.Printf("Unable to create temporary directory: %s\n", err.Error())
		os.Exit(1)
	}
	file, s, err := New(path.Join(dir, "app.log"), LevelDebug, "logtic")
	if err != nil {
		fmt.Printf("Unable to create new logtic instance: %s\n", err.Error())
		os.Exit(1)
	}
	defer file.Close()

	var wg sync.WaitGroup

	s.Debug("Start test")

	wg.Add(3)
	go func() {
		defer wg.Done()
		source := Connect("goroutine1")
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
		source := Connect("goroutine2")
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
		source := Connect("goroutine3")
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
	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		defer wg.Done()
		source := Connect("goroutine1")
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
		source := Connect("goroutine2")
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
		source := Connect("goroutine3")
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
	dir, err := ioutil.TempDir("", "logtic")
	if err != nil {
		fmt.Printf("Unable to create temporary directory: %s\n", err.Error())
		os.Exit(1)
	}
	file, s, err := New(path.Join(dir, "app.log"), LevelDebug, "logtic")
	if err != nil {
		fmt.Printf("Unable to create new logtic instance: %s\n", err.Error())
		os.Exit(1)
	}
	defer file.Close()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		i := 0
		for i < 100 {
			i++
			s.Debug("Count %d", i)
			s.Info("Count %d", i)
			s.Warn("Count %d", i)
			s.Error("Count %d", i)
		}
	}()
	time.Sleep(1 * time.Millisecond)
	if err := file.Rotate(); err != nil {
		t.Fatalf("Error rotating log file: %s", err.Error())
	}
	wg.Wait()

	date := time.Now().Format("2006-01-02")
	expectedPath := path.Join(dir, "app."+date+".log")
	if _, err := os.Stat(expectedPath); err != nil {
		t.Errorf("Expected rotated log file not found: '%s'", expectedPath)
	}

	expectedPath = path.Join(dir, "app.log")
	if _, err := os.Stat(expectedPath); err != nil {
		t.Errorf("Expected new log file not found: '%s'", expectedPath)
	}
}
