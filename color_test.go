package logtic_test

import (
	"bytes"
	"testing"

	"github.com/ecnepsnai/logtic"
)

func TestColor(t *testing.T) {
	Setup()

	b := &bytes.Buffer{}
	logtic.Log.Stdout = b
	logtic.Log.Stderr = b

	logtic.Log.Level = logtic.LevelWarn
	logtic.Log.Options.Color = false
	logtic.Log.Open()

	source := logtic.Log.Connect("example")
	source.Error("Warning")

	Setup()

	buf := b.Bytes()
	if buf[0] == 0x1b && buf[1] == 0x5b && buf[2] == 0x33 && buf[3] == 0x31 && buf[4] == 0x6d && buf[5] == 0x5b {
		t.Errorf("Printed output contains colour when it should not")
	}
}
