package log

import (
	"fmt"
	stdlog "log"
)

// F is the main log function.
var F = func(string, ...interface{}) {}

// Debugf prints debug log.
func Debugf(f string, v ...interface{}) {
	stdlog.SetFlags(stdlog.LstdFlags | stdlog.Lshortfile)
	stdlog.Output(2, fmt.Sprintf(f, v...))
}

// Print prints log.
func Print(v ...interface{}) {
	stdlog.Print(v...)
}

// Printf prints log.
func Printf(f string, v ...interface{}) {
	stdlog.Printf(f, v...)
}

// Fatal log and exit.
func Fatal(v ...interface{}) {
	stdlog.Fatal(v...)
}

// Fatalf log and exit.
func Fatalf(f string, v ...interface{}) {
	stdlog.Fatalf(f, v...)
}
