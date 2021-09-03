package logtic_test

import (
	"fmt"
	"os"
	"path"
	"testing"
	"time"

	"github.com/ecnepsnai/logtic"
)

func TestRotate(t *testing.T) {
	logtic.Log.Reset()

	dir := t.TempDir()

	logtic.Log.FilePath = path.Join(dir, "app.log")
	logtic.Log.Level = logtic.LevelDebug

	if err := logtic.Log.Open(); err != nil {
		t.Fatalf("Error opening log file: %s", err.Error())
	}

	s := logtic.Log.Connect("Test")
	i := 0
	for i < 5 {
		i++
		s.Debug("Count %d", i)
	}

	if err := logtic.Log.Rotate(); err != nil {
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

func TestRotateDuplicate(t *testing.T) {
	logtic.Log.Reset()

	dir := t.TempDir()

	logtic.Log.FilePath = path.Join(dir, "app.log")
	logtic.Log.Level = logtic.LevelDebug

	if err := logtic.Log.Open(); err != nil {
		t.Fatalf("Error opening log file: %s", err.Error())
	}

	s := logtic.Log.Connect("Test")

	i := 0
	for i < 5 {
		i++
		y := 0
		for y < 5 {
			y++
			s.Debug("i=%d y=%d", i, y)
		}

		if err := logtic.Log.Rotate(); err != nil {
			panic(err)
		}

		s.Info("Rotated log")
	}

	date := time.Now().Format("2006-01-02")
	expectedPaths := []string{
		path.Join(dir, "app.log."+date),
		path.Join(dir, "app.log."+date+"-1"),
		path.Join(dir, "app.log."+date+"-2"),
		path.Join(dir, "app.log."+date+"-3"),
		path.Join(dir, "app.log."+date+"-4"),
	}

	for _, expectedPath := range expectedPaths {
		if _, err := os.Stat(expectedPath); err != nil {
			t.Errorf("Expected rotated log file not found: '%s'", expectedPath)
		}
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

	for _, expectedPath := range expectedPaths {
		checkFileSize(expectedPath)
	}
	checkFileSize(currentPath)
}
