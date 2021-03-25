package logtic

import (
	"bytes"
	"os"
	"testing"
)

func TestColor(t *testing.T) {
	b := &bytes.Buffer{}

	Reset()
	Log.Level = LevelWarn
	Log.Color = false
	stdout = b
	stderr = b
	Open()

	source := Connect("example")
	source.Error("Warning")

	stdout = os.Stdout
	stderr = os.Stderr

	buf := b.Bytes()
	if buf[0] == 0x1b && buf[1] == 0x5b && buf[2] == 0x33 && buf[3] == 0x31 && buf[4] == 0x6d && buf[5] == 0x5b {
		t.Errorf("Printed output contains colour when it should not")
	}
}
