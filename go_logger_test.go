package logtic_test

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/ecnepsnai/logtic"
)

func TestGoLogger(t *testing.T) {
	Setup()

	b := &bytes.Buffer{}
	logtic.Log.Stdout = b
	logtic.Log.Stderr = b

	logtic.Log.Level = logtic.LevelError
	logtic.Log.Open()

	source := logtic.Log.Connect("example")
	l := source.GoLogger(logtic.LevelError)
	l.Print("Example")
	time.Sleep(100 * time.Millisecond)

	if !strings.Contains(b.String(), "example") {
		t.Errorf("Output did not contain expected message")
	}
}
