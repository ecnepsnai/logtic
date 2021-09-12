package logtic

import (
	"bytes"
	"os"
	"testing"
)

func TestEscapeCharacters(t *testing.T) {
	check := func(in, expect string) {
		result := escapeCharacters(in)
		if result != expect {
			t.Errorf("Incorrect result for escaped string. Got '%s' expected '%s'", result, expect)
		}
	}

	check("Hello\nWorld!", "Hello\\nWorld!")
	check("Hello\\nWorld!", "Hello\\nWorld!")
}

// Test that nothing is printed to the console when the logger is not opened
func TestLoggerNotOpened(t *testing.T) {
	Log.Reset()
	b := &bytes.Buffer{}
	stdout = b
	stderr = b

	source := Log.Connect("example")
	source.Debug("Debug")
	source.Info("Info")
	source.Warn("Warn")
	source.Error("Error")

	if b.Len() != 0 {
		t.Errorf("Data written to stdout/stderr that wasn't expected")
	}

	// Panic and Fatal messages are always printed to stderr
	defer func() {
		if r := recover(); r != nil {
			Log.Reset()
			stdout = os.Stdout
			stderr = os.Stderr

			if b.Len() == 0 {
				t.Errorf("Panic message not printed to stderr when expected")
			}
		}
	}()
	source.Panic("Panic!")
}
