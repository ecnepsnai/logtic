package logtic

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

// Source describes a source for log events
type Source struct {
	Name  string
	Level int
	dummy bool
}

func (s *Source) checkLevel(levelWanted int) bool {
	if s.Level >= 0 {
		return s.Level < levelWanted
	}
	return Log.Level < levelWanted
}

// Debug will log a debug message
func (s *Source) Debug(format string, a ...interface{}) {
	if s == nil || Log.file == nil || s.checkLevel(LevelDebug) {
		return
	}
	message := fmt.Sprintf(format, a...)
	fmt.Printf("%s %s\n", color.HiBlackString("[DEBUG]["+s.Name+"]"), message)
	Log.write("[DEBUG][" + s.Name + "] " + message)
}

// Info will log an informational message
func (s *Source) Info(format string, a ...interface{}) {
	if s == nil || Log.file == nil || s.checkLevel(LevelInfo) {
		return
	}
	message := fmt.Sprintf(format, a...)
	fmt.Printf("%s %s\n", color.BlueString("[INFO]["+s.Name+"]"), message)
	Log.write("[INFO][" + s.Name + "] " + message)
}

// Warn will log a warning message
func (s *Source) Warn(format string, a ...interface{}) {
	if s == nil || Log.file == nil || s.checkLevel(LevelWarn) {
		return
	}
	message := fmt.Sprintf(format, a...)
	fmt.Printf("%s %s\n", color.YellowString("[WARN]["+s.Name+"]"), message)
	Log.write("[WARN][" + s.Name + "] " + message)
}

// Error will log an error message. Errors are printed to stderr.
func (s *Source) Error(format string, a ...interface{}) {
	if s == nil || Log.file == nil || s.checkLevel(LevelError) {
		return
	}
	message := fmt.Sprintf(format, a...)
	fmt.Fprintf(os.Stderr, "%s %s\n", color.RedString("[ERROR]["+s.Name+"]"), message)
	Log.write("[ERROR][" + s.Name + "] " + message)
}

// Fatal will log a fatal error message and exit the application with status 1. Fatal messages are printed to stderr.
func (s *Source) Fatal(format string, a ...interface{}) {
	if s != nil && !s.dummy {
		message := fmt.Sprintf(format, a...)
		fmt.Fprintf(os.Stderr, "%s %s\n", color.RedString("[FATAL]["+s.Name+"]"), message)
		Log.write("[FATAL][" + s.Name + "] " + message)
	}
	os.Exit(1)
}

// Panic functions like source.Fatal() but panics rather than exits.
func (s *Source) Panic(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)
	if s != nil && !s.dummy {
		fmt.Fprintf(os.Stderr, "%s %s\n", color.RedString("[FATAL]["+s.Name+"]"), message)
		Log.write("[FATAL][" + s.Name + "] " + message)
	}
	panic(message)
}

// Write will call the matching write function for the given level, printing the provided message.
// For example:
//     source.Write(logtic.LevelDebug, "Hello world")
// is the same as:
//     source.Debug("Hello world")
func (s *Source) Write(level int, format string, a ...interface{}) {
	switch level {
	case LevelDebug:
		s.Debug(format, a...)
	case LevelInfo:
		s.Info(format, a...)
	case LevelWarn:
		s.Warn(format, a...)
	case LevelError:
		s.Error(format, a...)
	default:
		return
	}
}

func (l *Settings) write(message string) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.file.WriteString(time.Now().Format(time.RFC3339) + " " + message + "\n")
}
