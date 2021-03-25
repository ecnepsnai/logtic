package logtic_test

import (
	"os"

	"github.com/ecnepsnai/logtic"
)

// Common setup for file-based logging
func Example() {
	// Set the log file path and default log level (if desired)
	logtic.Log.FilePath = "./file.log"
	logtic.Log.Level = logtic.LevelInfo

	if err := logtic.Open(); err != nil {
		// There was an error opening the log file for writing
		panic(err)
	}

	log := logtic.Connect("MyApp")
	log.Warn("Print something %s", "COOL!")

	// Don't forget to close the log file when your application exits
	logtic.Close()
}

// This example shows how to prepare logtic for writing to a log file
func ExampleOpen() {
	// You must tell logtic where the log file is
	// before any events will be captured
	logtic.Log.FilePath = "./file.log"
	// The default level is Error, you can change that at any time
	logtic.Log.Level = logtic.LevelInfo

	if err := logtic.Open(); err != nil {
		// There was an error opening the log file for writing
		panic(err)
	}
}

// This example shows how to prepare logtic to only print to the console, without writing to a log file
func ExampleOpen_withoutFile() {
	// You don't have to specify os.DevNull, however it is good pratice to be explicit that it is
	// not being written to any file
	logtic.Log.FilePath = os.DevNull
	// The default level is Error, you can change that at any time
	logtic.Log.Level = logtic.LevelInfo

	if err := logtic.Open(); err != nil {
		// You probably should panic here, or exit the app
		// as failing to open a file to DevNull is a bad sign
		panic(err)
	}
}

// This example shows how to connect a new source to a logtic instance
func ExampleConnect() {
	// You can connect to logtic before a log file has been opened
	// however, any events will not be captured until logtic.Open()
	// has been called
	source1 := logtic.Connect("Source1")
	source2 := logtic.Connect("Source2")

	source1.Warn("Important warning")
	source2.Error("Something went wrong")
}

func ExampleSource_Debug() {
	log := logtic.Connect("Example")
	log.Debug("This is a %s message", "debug")
}

func ExampleSource_Info() {
	log := logtic.Connect("Example")
	log.Info("This is a %s message", "info")
}

func ExampleSource_Warn() {
	log := logtic.Connect("Example")
	log.Warn("This is a %s message", "warning")
}

func ExampleSource_Error() {
	log := logtic.Connect("Example")
	log.Error("This is a %s message", "error")
}

func ExampleSource_Fatal() {
	log := logtic.Connect("Example")
	log.Fatal("This is a %s message", "fatal")
}

func ExampleSource_Panic() {
	log := logtic.Connect("Example")
	log.Panic("This is a %s message", "fatal")
}

// This example shows how to trigger a log file rotation
func ExampleRotate() {
	logtic.Log.FilePath = "/path/to/log/file.log"

	if err := logtic.Open(); err != nil {
		panic(err)
	}

	if err := logtic.Rotate(); err != nil {
		// There was an error rotating the log
		// It's recommended that you panic or exit here, as logtic is now in an undefined state
		panic(err)
	}
}
