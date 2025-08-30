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
	logtic.Log.Open()

	source := logtic.Log.Connect("example")
	source.Error("Warning")

	Setup()

	buf := b.Bytes()
	expected := []byte{0x1b, 0x5b, 0x33, 0x31, 0x6d, 0x5b, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5d, 0x5b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5d, 0x1b, 0x5b, 0x30, 0x6d, 0x20, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x0a}

	if !bytes.Equal(buf, expected) {
		t.Fatalf("Unexpected output.\nExpected:\n\t%x\nGot:\n\t%x", expected, buf)
	}
}

func TestNoColor(t *testing.T) {
	Setup()

	b := &bytes.Buffer{}
	logtic.Log.Stdout = b
	logtic.Log.Stderr = b

	logtic.Log.Level = logtic.LevelWarn
	logtic.Log.Color = nil
	logtic.Log.Open()

	source := logtic.Log.Connect("example")
	source.Error("Warning")

	Setup()

	buf := b.Bytes()
	expected := []byte{0x5b, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5d, 0x5b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5d, 0x20, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x0a}

	if !bytes.Equal(buf, expected) {
		t.Fatalf("Unexpected output.\nExpected:\n\t%x\nGot:\n\t%x", expected, buf)
	}
}
