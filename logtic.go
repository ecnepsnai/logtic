// Package logtic is a yet another logging library for golang projects.
//
// The goal of logtic is to be as transparent and easy to use as possible, allowing applications and libraries to
// seamlessly log to a single file. Logtic can be used in libraries and won't cause any problems if the parent
// application isn't using logtic.
//
// Logtic supports multiple sources, which annotate the outputted log lines. It also supports defining a minimum
// desired log level, which can be changed at any time. Events printed to the terminal output support color-coded
// severities.
//
// Events can be printed as formatted strings, like with `fmt.Printf`, or can be parameterized events which can be
// easily parsed by log analysis tools such as Splunk.
//
// By default, logtic will only print to stdout and stderr, but when configured it can also write to a log file. Log
// files include the date-time for each event in RFC-3339 format.
//
// Logtic provides a default logging instance but also supports unique instances that can operate in parallel, writing
// to unique files and having unique settings.
//
// Log files can be rotated using the provided rotate method.
package logtic
