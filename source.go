package logtic

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/debug"
	"strings"
)

// Source describes a source for log events
type Source struct {
	Name     string
	Level    int
	instance *Logger
}

func (s *Source) formatMessage(format string, a ...interface{}) string {
	message := fmt.Sprintf(format, a...)
	if s != nil && s.instance != nil && s.instance.Options.EscapeCharacters {
		message = escapeCharacters(message)
	}
	return message
}

func (s *Source) write(message string) {
	s.instance.write(message)
}

func (s *Source) checkLevel(levelWanted int) bool {
	if s.Level >= 0 {
		return s.Level < levelWanted
	}
	return s.instance.Level < levelWanted
}

func (s *Source) stdout() io.Writer {
	if s == nil || s.instance == nil {
		return os.Stdout
	}
	return s.instance.Stdout
}

func (s *Source) stderr() io.Writer {
	if s == nil || s.instance == nil {
		return os.Stderr
	}
	return s.instance.Stderr
}

// Debug will log a debug formatted message.
func (s *Source) Debug(format string, a ...interface{}) {
	defer panicRecover()
	if s == nil || s.instance == nil || !s.instance.opened || s.checkLevel(LevelDebug) {
		return
	}
	message := s.formatMessage(format, a...)
	fmt.Fprintf(s.stdout(), "%s %s\n", colorHiBlackString("[DEBUG]["+s.Name+"]"), message)
	s.write("[DEBUG][" + s.Name + "] " + message)
}

// Info will log an informational formatted message.
func (s *Source) Info(format string, a ...interface{}) {
	defer panicRecover()
	if s == nil || s.instance == nil || !s.instance.opened || s.checkLevel(LevelInfo) {
		return
	}
	message := s.formatMessage(format, a...)
	fmt.Fprintf(s.stdout(), "%s %s\n", colorBlueString("[INFO]["+s.Name+"]"), message)
	s.write("[INFO][" + s.Name + "] " + message)
}

// Warn will log a warning formatted message.
func (s *Source) Warn(format string, a ...interface{}) {
	defer panicRecover()
	if s == nil || s.instance == nil || !s.instance.opened || s.checkLevel(LevelWarn) {
		return
	}
	message := s.formatMessage(format, a...)
	fmt.Fprintf(s.stdout(), "%s %s\n", colorYellowString("[WARN]["+s.Name+"]"), message)
	s.write("[WARN][" + s.Name + "] " + message)
}

// Error will log an error formatted message. Errors are printed to stderr.
func (s *Source) Error(format string, a ...interface{}) {
	defer panicRecover()
	if s == nil || s.instance == nil || !s.instance.opened || s.checkLevel(LevelError) {
		return
	}
	message := s.formatMessage(format, a...)
	fmt.Fprintf(s.stderr(), "%s %s\n", colorRedString("[ERROR]["+s.Name+"]"), message)
	s.write("[ERROR][" + s.Name + "] " + message)
}

// Fatal will log a fatal formatted error message and exit the application with status 1.
// Fatal messages are printed to stderr.
func (s *Source) Fatal(format string, a ...interface{}) {
	message := s.formatMessage(format, a...)
	fmt.Fprintf(s.stderr(), "%s %s\n", colorRedString("[FATAL]["+s.Name+"]"), message)
	s.write("[FATAL][" + s.Name + "] " + message)
	os.Exit(1)
}

// Panic functions like source.Fatal() but panics rather than exits.
func (s *Source) Panic(format string, a ...interface{}) {
	message := s.formatMessage(format, a...)
	fmt.Fprintf(s.stderr(), "%s %s\n", colorRedString("[FATAL]["+s.Name+"]"), message)
	s.write("[FATAL][" + s.Name + "] " + message)
	panic(message)
}

// Write will call the matching write function for the given level, printing the provided message.
// For example:
//
//	source.Write(logtic.LevelDebug, "Hello world")
//
// is the same as:
//
//	source.Debug("Hello world")
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

func escapeCharacters(message string) string {
	message = strings.ReplaceAll(message, "\a", "\\a")
	message = strings.ReplaceAll(message, "\b", "\\b")
	message = strings.ReplaceAll(message, "\t", "\\t")
	message = strings.ReplaceAll(message, "\n", "\\n")
	message = strings.ReplaceAll(message, "\f", "\\f")
	message = strings.ReplaceAll(message, "\r", "\\r")
	message = strings.ReplaceAll(message, "\v", "\\v")
	return message
}

func panicRecover() {
	if r := recover(); r != nil {
		fmt.Fprintf(os.Stderr, "logtic: recovered from panic writing event. stack to follow.\n")
		debug.PrintStack()
	}
}

// GoLogger returns a logger that acts as a proxy between the go/log package and logtic.
// Printf events sent to this logger will be forwarded to this source with the given level.
func (s *Source) GoLogger(level int) *log.Logger {
	b := &bytes.Buffer{}

	go func() {
		for {
			message, err := b.ReadString('\n')
			if err != nil {
				break
			}
			s.Write(level, message[0:len(message)-1])
		}
	}()

	return log.New(b, "", 0)
}
