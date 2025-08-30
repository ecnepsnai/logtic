package logtic_test

import (
	"fmt"
	"os"
	"time"

	"github.com/ecnepsnai/logtic"
)

// Common setup for file-based logging
func Example() {
	// Set the log file path and default log level (if desired)
	logtic.Log.FilePath = "./file.log"
	logtic.Log.Level = logtic.LevelInfo

	if err := logtic.Log.Open(); err != nil {
		// There was an error opening the log file for writing
		panic(err)
	}

	log := logtic.Log.Connect("MyApp")
	log.Warn("Print something %s", "COOL!")

	// Don't forget to close the log file when your application exits
	logtic.Log.Close()
}

// This example shows how to set up a unique logging instance, separate form the default instance
func ExampleLogger_Open() {
	logger := logtic.New()

	// You must tell logtic where the log file is
	// before any events will be captured
	logger.FilePath = "./file.log"
	// The default level is Error, you can change that at any time
	logger.Level = logtic.LevelInfo

	if err := logger.Open(); err != nil {
		// There was an error opening the log file for writing
		panic(err)
	}
}

// This example shows how to prepare logtic to only print to the console, without writing to a log file
func ExampleLogger_Open_withoutFile() {
	logger := logtic.New()

	// You don't have to specify os.DevNull, however it is good pratice to be explicit that it is
	// not being written to any file
	logger.FilePath = os.DevNull
	// The default level is Error, you can change that at any time
	logger.Level = logtic.LevelInfo

	if err := logger.Open(); err != nil {
		// You probably should panic here, or exit the app
		// as failing to open a file to DevNull is a bad sign
		panic(err)
	}
}

// This example shows how to connect a new source to a logtic instance
func ExampleLogger_Connect() {
	// You can connect to logtic before a log file has been opened
	// however, any events will not be captured until logtic.Log.Open()
	// has been called
	source1 := logtic.Log.Connect("Source1")
	source2 := logtic.Log.Connect("Source2")

	source1.Warn("Important warning")
	source2.Error("Something went wrong")
}

func ExampleSource_Debug() {
	log := logtic.Log.Connect("Example")
	log.Debug("This is a %s message", "debug")
	// Terminal output: [DEBUG][Example] This is a debug message
	// File output: 2021-03-15T21:43:34-07:00 [DEBUG][Example] This is a debug message
}

func ExampleSource_Info() {
	log := logtic.Log.Connect("Example")
	log.Info("This is a %s message", "info")
	// Terminal output: [INFO][Example] This is a info message
	// File output: 2021-03-15T21:43:34-07:00 [INFO][Example] This is a info message
}

func ExampleSource_Warn() {
	log := logtic.Log.Connect("Example")
	log.Warn("This is a %s message", "warning")
	// Terminal output: [WARN][Example] This is a warning message
	// File output: 2021-03-15T21:43:34-07:00 [WARN][Example] This is a warning message
}

func ExampleSource_Error() {
	log := logtic.Log.Connect("Example")
	log.Error("This is a %s message", "error")
	// Terminal output: [ERROR][Example] This is a error message
	// File output: 2021-03-15T21:43:34-07:00 [ERROR][Example] This is a error message
}

func ExampleSource_Fatal() {
	log := logtic.Log.Connect("Example")
	log.Fatal("This is a %s message", "fatal")
	// Terminal output: [FATAL][Example] This is a fatal message
	// File output: 2021-03-15T21:43:34-07:00 [FATAL][Example] This is a fatal message
}

func ExampleSource_Panic() {
	log := logtic.Log.Connect("Example")
	log.Panic("This is a %s message", "fatal")
	// Terminal output: [FATAL][Example] This is a fatal message
	// File output: 2021-03-15T21:43:34-07:00 [FATAL][Example] This is a fatal message
}

func ExampleSource_PDebug() {
	log := logtic.Log.Connect("Example")
	log.PDebug("Debug event", map[string]any{
		"param1": "string",
		"param2": 123,
	})
	// Terminal output: [DEBUG][Example] Debug event: param1='string' param2=123
	// File output: 2021-03-15T21:43:34-07:00 [DEBUG][Example] Debug event: param1='string' param2=123
}

func ExampleSource_PInfo() {
	log := logtic.Log.Connect("Example")
	log.PInfo("Info event", map[string]any{
		"param1": "string",
		"param2": 123,
	})
	// Terminal output: [INFO][Example] Info event: param1='string' param2=123
	// File output: 2021-03-15T21:43:34-07:00 [INFO][Example] Info event: param1='string' param2=123
}

func ExampleSource_PWarn() {
	log := logtic.Log.Connect("Example")
	log.PWarn("Warning event", map[string]any{
		"param1": "string",
		"param2": 123,
	})
	// Terminal output: [WARN][Example] Warning event: param1='string' param2=123
	// File output: 2021-03-15T21:43:34-07:00 [WARN][Example] Warning event: param1='string' param2=123
}

func ExampleSource_PError() {
	log := logtic.Log.Connect("Example")
	log.PError("Error event", map[string]any{
		"param1": "string",
		"param2": 123,
	})
	// Terminal output: [ERROR][Example] Error event: param1='string' param2=123
	// File output: 2021-03-15T21:43:34-07:00 [ERROR][Example] Error event: param1='string' param2=123
}

func ExampleSource_PFatal() {
	log := logtic.Log.Connect("Example")
	log.PFatal("Fatal event", map[string]any{
		"param1": "string",
		"param2": 123,
	})
	// Terminal output: [FATAL][Example] Fatal event: param1='string' param2=123
	// File output: 2021-03-15T21:43:34-07:00 [FATAL][Example] Fatal event: param1='string' param2=123
}

func ExampleSource_PPanic() {
	log := logtic.Log.Connect("Example")
	log.PPanic("Panic event", map[string]any{
		"param1": "string",
		"param2": 123,
	})
	// Terminal output: [FATAL][Example] Panic event: param1='string' param2=123
	// File output: 2021-03-15T21:43:34-07:00 [FATAL][Example] Panic event: param1='string' param2=123
}

// This example shows how to trigger a log file rotation
func ExampleLogger_RotateDate() {
	logtic.Log.FilePath = "/path/to/log/file.log"

	if err := logtic.Log.Open(); err != nil {
		panic(err)
	}

	if err := logtic.Log.RotateDate(); err != nil {
		// There was an error rotating the log
		// It's recommended that you panic or exit here, as logtic is now in an undefined state
		panic(err)
	}
}

func ExampleStringFromParameters() {
	fmt.Println(logtic.StringFromParameters(map[string]any{
		"hello":            "world!",
		"meaning_of_life":  42,
		"pie":              3.12,
		"does_golang_rock": true,
		"secret_sauce":     []byte("mayo ketchup sriracha"),
		"unix_epoch":       time.Unix(0, 0).UTC(),
		"prime_numbers":    []int{2, 3, 5, 7, 11},
	}))
	// output: does_golang_rock='true' hello='world!' meaning_of_life=42 pie=3.120000 prime_numbers='[2 3 5 7 11]' secret_sauce=6d61796f206b657463687570207372697261636861 unix_epoch='1970-01-01T00:00:00Z'
}

func ExampleFormatBytesB() {
	fmt.Println(logtic.FormatBytesB(10485760))
	// output: 10.0 MiB
}

func ExampleFormatBytesD() {
	fmt.Println(logtic.FormatBytesD(10000000))
	// output: 10.0 MB
}
