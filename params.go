package logtic

import (
	"fmt"
	"reflect"
	"sort"
	"time"
)

// StringFromParameters return a key=value string for the given parameters. Depending on the type of the parameter
// value, it may be wrapped in single quotes. Byte slices are represented as hexadecimal strings. Parameters are always
// alphabetically sorted in the outputted string.
func StringFromParameters(parameters map[string]any) string {
	out := ""
	last := len(parameters) - 1
	i := 0
	keys := make([]string, len(parameters))
	for k := range parameters {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for i, k := range keys {
		v := parameters[k]
		t := reflect.TypeOf(v)
		out += k + "="
		switch t.Kind() {
		case reflect.String:
			if t.AssignableTo(reflect.TypeOf("")) {
				out += "'" + v.(string) + "'"
			} else {
				out += fmt.Sprintf("'%v'", v)
			}
		case reflect.Int,
			reflect.Int8,
			reflect.Int16,
			reflect.Int32,
			reflect.Int64,
			reflect.Uint,
			reflect.Uint8,
			reflect.Uint16,
			reflect.Uint32,
			reflect.Uint64:
			out += fmt.Sprintf("%d", v)
		case reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
			out += fmt.Sprintf("%f", v)
		case reflect.Slice, reflect.Array:
			if b, isBytes := v.([]byte); isBytes {
				out += fmt.Sprintf("%x", b)
			} else {
				out += fmt.Sprintf("'%v'", v)
			}
		case reflect.Struct:
			if t, isTime := v.(time.Time); isTime {
				out += "'" + t.Format(time.RFC3339) + "'"
			} else {
				out += fmt.Sprintf("'%v'", v)
			}
		default:
			out += fmt.Sprintf("'%v'", v)
		}
		if i != last {
			out += " "
		}
	}
	return out
}

// FormatBytesB takes in a number of bytes and returns a human readable string with binary units (up-to Exbibyte)
func FormatBytesB(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}

// FormatBytesB takes in a number of bytes and returns a human readable string with decimal units (up-to Exabyte)
func FormatBytesD(b uint64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}

// PDebug will log a debug parameterized message.
// Parameterized messages are formatted as key=value strings. Depending on the type of the parameter value, it may
// be wrapped in single quotes. Byte slices are represented as hexadecimal strings. Parameters are always alphabetically
// sorted in the outputted string.
func (s *Source) PDebug(event string, parameters map[string]any) {
	s.Debug("%s: %s", event, StringFromParameters(parameters))
}

// PInfo will log an informational parameterized message.
// Parameterized messages are formatted as key=value strings. Depending on the type of the parameter value, it may
// be wrapped in single quotes. Byte slices are represented as hexadecimal strings. Parameters are always alphabetically
// sorted in the outputted string.
func (s *Source) PInfo(event string, parameters map[string]any) {
	s.Info("%s: %s", event, StringFromParameters(parameters))
}

// PWarn will log a warning parameterized message.
// Parameterized messages are formatted as key=value strings. Depending on the type of the parameter value, it may
// be wrapped in single quotes. Byte slices are represented as hexadecimal strings. Parameters are always alphabetically
// sorted in the outputted string.
func (s *Source) PWarn(event string, parameters map[string]any) {
	s.Warn("%s: %s", event, StringFromParameters(parameters))
}

// PError will log an error parameterized message. Errors are printed to stderr.
// Parameterized messages are formatted as key=value strings. Depending on the type of the parameter value, it may
// be wrapped in single quotes. Byte slices are represented as hexadecimal strings. Parameters are always alphabetically
// sorted in the outputted string.
func (s *Source) PError(event string, parameters map[string]any) {
	s.Error("%s: %s", event, StringFromParameters(parameters))
}

// PFatal will log a fatal parameterized error message and exit the application with status 1.
// Fatal messages are printed to stderr.
// Parameterized messages are formatted as key=value strings. Depending on the type of the parameter value, it may
// be wrapped in single quotes. Byte slices are represented as hexadecimal strings. Parameters are always alphabetically
// sorted in the outputted string.
func (s *Source) PFatal(event string, parameters map[string]any) {
	s.Fatal("%s: %s", event, StringFromParameters(parameters))
}

// PPanic functions like source.PFatal() but panics rather than exits.
// Parameterized messages are formatted as key=value strings. Depending on the type of the parameter value, it may
// be wrapped in single quotes. Byte slices are represented as hexadecimal strings. Parameters are always alphabetically
// sorted in the outputted string.
func (s *Source) PPanic(event string, parameters map[string]any) {
	s.Panic("%s: %s", event, StringFromParameters(parameters))
}

// PWrite will call the matching write function for the given level, printing the provided message.
// For example:
//
//	source.PWrite(logtic.LevelDebug, "My Event", map[string]any{"key": "value"})
//
// is the same as:
//
//	source.PDebug("My Event", map[string]any{"key": "value"})
func (s *Source) PWrite(level LogLevel, event string, parameters map[string]any) {
	switch level {
	case LevelDebug:
		s.PDebug(event, parameters)
	case LevelInfo:
		s.PInfo(event, parameters)
	case LevelWarn:
		s.PWarn(event, parameters)
	case LevelError:
		s.PError(event, parameters)
	default:
		return
	}
}
