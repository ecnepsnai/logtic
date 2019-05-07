package logtic

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

// Source describes a source for log events
type Source struct {
	dummy     bool
	className string
}

func dummySource() Source {
	return Source{
		dummy: true,
	}
}

// Debug log debug messages
func (s *Source) Debug(format string, a ...interface{}) {
	if s == nil || s.dummy == true || Log.Level < LevelDebug {
		return
	}
	message := fmt.Sprintf(format, a...)
	fmt.Printf("%s %s\n", color.HiBlackString("[DEBUG]["+s.className+"]"), message)
	Log.write("[DEBUG][" + s.className + "] " + message)
}

// Info log information messages
func (s *Source) Info(format string, a ...interface{}) {
	if s == nil || s.dummy == true || Log.Level < LevelInfo {
		return
	}
	message := fmt.Sprintf(format, a...)
	fmt.Printf("%s %s\n", color.BlueString("[INFO]["+s.className+"]"), message)
	Log.write("[INFO][" + s.className + "] " + message)
}

// Warn log warning messages
func (s *Source) Warn(format string, a ...interface{}) {
	if s == nil || s.dummy == true || Log.Level < LevelWarn {
		return
	}
	message := fmt.Sprintf(format, a...)
	fmt.Printf("%s %s\n", color.YellowString("[WARN]["+s.className+"]"), message)
	Log.write("[WARN][" + s.className + "] " + message)
}

// Error log error messages
func (s *Source) Error(format string, a ...interface{}) {
	if s == nil || s.dummy == true || Log.Level < LevelError {
		return
	}
	message := fmt.Sprintf(format, a...)
	fmt.Fprintf(os.Stderr, "%s %s\n", color.RedString("[ERROR]["+s.className+"]"), message)
	Log.write("[ERROR][" + s.className + "] " + message)
}

// Fatal log a fatal error then exit the application
func (s *Source) Fatal(format string, a ...interface{}) {
	if s != nil && !s.dummy {
		message := fmt.Sprintf(format, a...)
		fmt.Fprintf(os.Stderr, "%s %s\n", color.RedString("[FATAL]["+s.className+"]"), message)
		Log.write("[FATAL][" + s.className + "] " + message)
	}
	os.Exit(1)
}

// Panic log a fatal error then panic
func (s *Source) Panic(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)
	if s != nil && !s.dummy {
		fmt.Fprintf(os.Stderr, "%s %s\n", color.RedString("[FATAL]["+s.className+"]"), message)
		Log.write("[FATAL][" + s.className + "] " + message)
	}
	panic(message)
}

func (l *Settings) write(message string) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.file.WriteString(time.Now().Format(time.RFC3339) + " " + message + "\n")
}
