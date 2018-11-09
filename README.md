# logtic

Logtic is a (another) logging library for golang projects

Logtic is meant for large applications that contain multiple libraries
that all need to write to a single log file.
Logtic is transparent in that it can be included in your libraries and attach
to any log file if the parent application is using logtic, otherwise it
just does nothing.

# Usage

## In your app

```golang
file, source, err := logtic.New("app.log", logtic.LevelWarn, "myAppName")
defer file.Close()
source.Info("Print something %s", "COOL")
```

## In your library

```golang
source := logtic.Connect("myLibrary")
source.Info("Print something %s", "COOL")
```