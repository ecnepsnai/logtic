# logtic

[![Go Report Card](https://goreportcard.com/badge/github.com/ecnepsnai/logtic?style=flat-square)](https://goreportcard.com/report/github.com/ecnepsnai/logtic)
[![Godoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/ecnepsnai/logtic)
[![Releases](https://img.shields.io/github/release/ecnepsnai/logtic/all.svg?style=flat-square)](https://github.com/ecnepsnai/logtic/releases)
[![LICENSE](https://img.shields.io/github/license/ecnepsnai/logtic.svg?style=flat-square)](https://github.com/ecnepsnai/logtic/blob/master/LICENSE)

Logtic is a (another) logging library for golang projects

Logtic is meant for large applications that contain multiple libraries
that all need to write to a single log file.
Logtic is transparent in that it can be included in your libraries and attach
to any log file if the parent application is using logtic, otherwise it
just does nothing.
The overall goal of logtic is that "it just works", meaning there should be little
effort required to get it working the correct way.

# Setup

In your application code (for example, the main package), configure Logtic in the init method.

```golang
func init() {
    // Tell logtic the path to the log file
    logtic.Log.FilePath = "./app.log"
    // Specify the log level
    logtic.Log.Level = logtic.LevelDebug
    // Open the log file, this will create it if it does not exist
    if err := logtic.Open(); err != nil {
        panic(err)
    }
}
```

You only have to do this once.

# Usage

Logtic uses "sources", which can be thought of like Classes in traditional OOP. Log entries are annotated with the
name.

For example, if you have two unique goroutines that you want to easily identify in logs, you can give each routine
a source with a unique name.

```golang
log := logtic.Connect("MyCoolApplication")
log.Info("Print something %s", "COOL!")
```
