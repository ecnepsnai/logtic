package logtic_test

import (
	"os"
	"path"
	"regexp"
	"testing"
	"time"

	"github.com/ecnepsnai/logtic"
)

func TestSourceParameters(t *testing.T) {
	Setup()

	logPath := path.Join(t.TempDir(), "logtic.log")

	logtic.Log.FilePath = logPath
	logtic.Log.Level = logtic.LevelDebug

	if err := logtic.Log.Open(); err != nil {
		t.Fatalf("Error opening log file: %s", err.Error())
	}

	parameters := map[string]any{
		"string": "hello, world!",
		"int":    123,
		"float":  3.14,
		"bool":   true,
		"bytes":  []byte("Hello, world!"),
		"slice":  []int{1, 2, 3},
		"nil":    nil,
	}
	source := logtic.Log.Connect("test")
	source.PWrite(logtic.LevelDebug, "Event", parameters)
	source.PWrite(logtic.LevelInfo, "Event", parameters)
	source.PWrite(logtic.LevelWarn, "Event", parameters)
	source.PWrite(logtic.LevelError, "Event", parameters)
	pattern := regexp.MustCompile(`[0-9\-:TZ]+ \[(DEBUG|INFO|WARN|ERROR)\]\[test\] Event: bool='true' bytes=48656c6c6f2c20776f726c6421 float=3.140000 int=123 nil=nil slice='\[1 2 3\]' string='hello, world!'`)

	logtic.Log.Close()

	logFileData, err := os.ReadFile(logPath)
	if err != nil {
		panic(err)
	}

	if !pattern.Match(logFileData) {
		t.Errorf("Log file does not contain expected log line for Debug message")
	}
}

type ExampleType string

func TestStringFromParameters(t *testing.T) {
	test := func(in interface{}, expected string) {
		out := logtic.StringFromParameters(map[string]any{
			"key": in,
		})
		result := out[4:]
		if result != expected {
			t.Errorf("Unexpected result for StringFromParameters. For %v Expected \"%s\" got \"%s\"", in, expected, result)
		}
	}

	test("hello", "'hello'")
	test(123, "123")
	test(true, "'true'")
	test([]byte("hello"), "68656c6c6f")
	test([]int{1, 2, 3}, "'[1 2 3]'")
	test(time.Unix(0, 0).UTC(), "'1970-01-01T00:00:00Z'")
	test(map[string]string{"hello": "world"}, "'map[hello:world]'")
	test(struct{ hello string }{hello: "world"}, "'{world}'")
	test(3.14, "3.140000")
	test(ExampleType("hello"), "'hello'")
}

func TestFormatBytesB(t *testing.T) {
	test := func(in uint64, expected string) {
		result := logtic.FormatBytesB(in)
		if result != expected {
			t.Errorf("Unexpected result from FormatBytesB. For %d Expected '%s' got '%s'", in, expected, result)
		}
	}

	test(100, "100 B")
	test(1024, "1.0 KiB")
	test(10240, "10.0 KiB")
	test(102400, "100.0 KiB")
	test(1024000, "1000.0 KiB")
	test(438143210, "417.8 MiB")
	test(57435943275, "53.5 GiB")
	test(2482587438925, "2.3 TiB")
	test(957183938585752, "870.6 TiB")
	test(65718393858575225, "58.4 PiB")
}
func TestFormatBytesD(t *testing.T) {
	test := func(in uint64, expected string) {
		result := logtic.FormatBytesD(in)
		if result != expected {
			t.Errorf("Unexpected result from FormatBytesD. For %d Expected '%s' got '%s'", in, expected, result)
		}
	}

	test(100, "100 B")
	test(1000, "1.0 KB")
	test(10000, "10.0 KB")
	test(100000, "100.0 KB")
	test(1000000, "1.0 MB")
	test(417800000, "417.8 MB")
	test(53500000000, "53.5 GB")
	test(2300000000000, "2.3 TB")
	test(870600000000000, "870.6 TB")
	test(58400000000000000, "58.4 PB")
}
