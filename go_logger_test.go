package logtic_test

import (
	"bytes"
	"strings"
	"testing"

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

	if !strings.Contains(b.String(), "example") {
		t.Errorf("Output did not contain expected message. Expected to see 'example' in '%s'", b.String())
	}
}
