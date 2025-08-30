package logtic

import (
	"log"
)

// GoLogger returns a logger that acts as a proxy between the go/log package and logtic.
// Printf events sent to this logger will be forwarded to this source with the given level.
func (s *Source) GoLogger(level LogLevel) *log.Logger {
	return log.New(&tGoLogger{
		level:  level,
		source: s,
	}, "", 0)
}

type tGoLogger struct {
	level  LogLevel
	source *Source
}

func (t *tGoLogger) Write(p []byte) (n int, err error) {
	var message string
	if p[len(p)-1] == '\n' {
		message = string(p[len(p)-1])
	} else {
		message = string(p)
	}
	t.source.Write(t.level, message)

	return len(p), nil
}
