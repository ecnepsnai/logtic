# logtic

[![Go Report Card](https://goreportcard.com/badge/github.com/ecnepsnai/logtic?style=flat-square)](https://goreportcard.com/report/github.com/ecnepsnai/logtic)
[![Godoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/ecnepsnai/logtic)
[![Releases](https://img.shields.io/github/release/ecnepsnai/logtic/all.svg?style=flat-square)](https://github.com/ecnepsnai/logtic/releases)
[![LICENSE](https://img.shields.io/github/license/ecnepsnai/logtic.svg?style=flat-square)](https://github.com/ecnepsnai/logtic/blob/master/LICENSE)

Package logtic is a (another) logging library for golang projects.

The goal of logtic is to be as transparent and easy to use as possible, allowing applications and libraries to
seamlessly log to a single file. Logtic can be used in libraries and won't cause any problems if the parent
application isn't using logtic.

Logtic supports multiple sources, which annotate the outputted log lines. It also supports defining a minimum
desired log level, which can be changed at any time.

By default, logtic will only print to stdout and stderr, but when configured it can also write to a log file.
Logtic can also rotate these log files out by invoking the logtic.Rotate() method. Log files include the date-time
for each line in RFC-3339 format.

# Usage & Examples

**For more examples please refer to the official [documentation](https://pkg.go.dev/github.com/ecnepsnai/logtic)**

## Basic Log File

```go
var log *logtic.Source

func init() {
    logtic.Log.FilePath = "./app.log"
    logtic.Log.Level = logtic.LevelInfo
}

func main() {
    if err := logtic.Open(); err != nil {
        // There was an error opening the log file for writing
        panic(err)
    }
    log = logtic.Connect("MyApp")
    defer logtic.Close()
}
```

## Console-Only (No Log File)

```go
var log *logtic.Source

func main() {
    // By default, calling logtic.Open without changing anything
    // will open a console-only log file
    logtic.Open()
    log = logtic.Connect("MyApp")
}
```

## Writing Events

Logtic uses "sources", which can be thought of like Classes in traditional OOP. Log entries are annotated with the
name.

For example, if you have two unique goroutines that you want to easily identify in logs, you can give each routine
a source with a unique name.

```go
log := logtic.Connect("MyCoolApplication")
log.Info("Print something %s", "COOL!")
// Output: [INFO][MyCoolApplication] Print something COOL!
```
