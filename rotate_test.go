package logtic_test

import (
	"bytes"
	"os"
	"path"
	"testing"
	"time"

	"github.com/ecnepsnai/logtic"
)

func fileIsGreaterThan1Byte(inPath string, t *testing.T) {
	info, err := os.Stat(inPath)
	if err != nil {
		t.Fatalf("Error stating rotated log file: %s", err.Error())
	}

	t.Logf("%s -> %d", inPath, info.Size())

	if info.Size() == 0 {
		t.Errorf("Rotated log file is empty")
	}
}

func TestRotate(t *testing.T) {
	Setup()

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

	fileIsGreaterThan1Byte(expectedPath, t)
	fileIsGreaterThan1Byte(currentPath, t)
}

func TestRotateDuplicate(t *testing.T) {
	Setup()

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

	for _, expectedPath := range expectedPaths {
		fileIsGreaterThan1Byte(expectedPath, t)
	}
	fileIsGreaterThan1Byte(currentPath, t)
}

func TestRotateGZip(t *testing.T) {
	Setup()

	dir := t.TempDir()

	logtic.Log.FilePath = path.Join(dir, "app.log")
	logtic.Log.Level = logtic.LevelDebug
	logtic.Log.Options.GZipRotatedLogs = true

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
	expectedPath := path.Join(dir, "app.log."+date+".gz")
	if _, err := os.Stat(expectedPath); err != nil {
		t.Errorf("Expected rotated log file not found: '%s'", expectedPath)
	}

	currentPath := path.Join(dir, "app.log")
	if _, err := os.Stat(currentPath); err != nil {
		t.Errorf("Expected new log file not found: '%s'", currentPath)
	}

	fileIsGreaterThan1Byte(expectedPath, t)
	fileIsGreaterThan1Byte(currentPath, t)

	gzipMagicNumber := []byte{0x1F, 0x8B}
	magicNumber := make([]byte, 2)
	f, err := os.OpenFile(expectedPath, os.O_RDONLY, 0644)
	if err != nil {
		t.Fatalf("Error reading rotated log file: %s", err.Error())
	}
	defer f.Close()
	f.Read(magicNumber)
	if !bytes.Equal(gzipMagicNumber, magicNumber) {
		t.Errorf("Incorrect file signature for rotated log file. Got '%x' expected '%x'", magicNumber, gzipMagicNumber)
	}
}
