// Package llog implements level logging on top of log package.
package llog

import (
	"fmt"
	"io"
	"log"
	"os"
)

type severity uint8

const (
	DEBUG severity = iota
	INFO
	WARNING
	ERROR
	FATAL
	PANIC
)

var severityName = []string{
	DEBUG:   "DEBUG",
	INFO:    "INFO",
	WARNING: "WARNING",
	ERROR:   "ERROR",
	FATAL:   "FATAL",
	PANIC:   "PANIC",
}

type Logger struct {
	threshold      severity
	severityPrefix bool
	stdlog         *log.Logger
}

// New creates a new logger
func New(threshold severity, output io.Writer) *Logger {
	return &Logger{threshold: threshold, stdlog: log.New(output, "", log.LstdFlags)}
}

var std = New(INFO, os.Stderr)

func (l *Logger) print(s severity, v ...interface{}) {
	if s < l.threshold {
		return
	}
	if l.severityPrefix {
		v = append([]interface{}{fmt.Sprintf("%s - ", severityName[s])}, v...)
	}
	switch s {
	case FATAL:
		l.stdlog.Fatal(v...)
	case PANIC:
		l.stdlog.Panic(v...)
	default:
		l.stdlog.Print(v...)
	}
}

func (l *Logger) printf(s severity, format string, v ...interface{}) {
	if s < l.threshold {
		return
	}
	if l.severityPrefix {
		format = fmt.Sprintf("%s - %s", severityName[s], format)
	}
	switch s {
	case FATAL:
		l.stdlog.Fatalf(format, v...)
	case PANIC:
		l.stdlog.Panicf(format, v...)
	default:
		l.stdlog.Printf(format, v...)
	}
}

func (l *Logger) println(s severity, v ...interface{}) {
	if s < l.threshold {
		return
	}
	if l.severityPrefix {
		v = append([]interface{}{fmt.Sprintf("%s -", severityName[s])}, v...)
	}
	switch s {
	case FATAL:
		l.stdlog.Fatalln(v...)
	case PANIC:
		l.stdlog.Panicln(v...)
	default:
		l.stdlog.Println(v...)
	}
}

// SetPrefix sets the output prefix for the logger.
func (l *Logger) SetPrefix(prefix string) {
	l.stdlog.SetPrefix(prefix)
}

// Threshold returns the current log threshold
func (l *Logger) Threshold() severity {
	return l.threshold
}

// SetThreshold sets a new log threshold.
// s denotes the lowest level to log
func (l *Logger) SetThreshold(s severity) {
	l.threshold = s
}

// Pdate enables/disables date output in prefix.
// Default: enabled
func (l *Logger) Pdate(set bool) {
	if set {
		l.stdlog.SetFlags(l.stdlog.Flags() | log.Ldate)
	} else {
		l.stdlog.SetFlags(l.stdlog.Flags() & ^log.Ldate)
	}
}

// Ptime enables/disables time output in prefix.
// Default: enabled
func (l *Logger) Ptime(set bool) {
	if set {
		l.stdlog.SetFlags(l.stdlog.Flags() | log.Ltime)
	} else {
		l.stdlog.SetFlags(l.stdlog.Flags() & ^log.Ltime)
	}
}

// Pmicroseconds enables/disables microseconds output in prefix.
// Default: disabled
func (l *Logger) Pmicroseconds(set bool) {
	if set {
		l.stdlog.SetFlags(l.stdlog.Flags() | log.Lmicroseconds)
	} else {
		l.stdlog.SetFlags(l.stdlog.Flags() & ^log.Lmicroseconds)
	}
}

// Pshortfile enables/disables short filename output in prefix.
// Enabling Pshortfile, disables Plongfile.
// Default: disabled
func (l *Logger) Pshortfile(set bool) {
	if set {
		l.stdlog.SetFlags(l.stdlog.Flags() & ^log.Llongfile)
		l.stdlog.SetFlags(l.stdlog.Flags() | log.Lshortfile)
	} else {
		l.stdlog.SetFlags(l.stdlog.Flags() & ^log.Lshortfile)
	}
}

// Plongfile enables/disables long (full path) filename output in prefix.
// Enabling Plongfile, disables Pshortfile.
// Default: disabled
func (l *Logger) Plongfile(set bool) {
	if set {
		l.stdlog.SetFlags(l.stdlog.Flags() & ^log.Lshortfile)
		l.stdlog.SetFlags(l.stdlog.Flags() | log.Llongfile)
	} else {
		l.stdlog.SetFlags(l.stdlog.Flags() & ^log.Llongfile)
	}
}

// Pseverity enables/disables severity output in prefix.
// Default: disabled
func (l *Logger) Pseverity(set bool) {
	l.severityPrefix = set
}

// Debug writes debug message with log.Print
func (l *Logger) Debug(v ...interface{}) {
	l.print(DEBUG, v...)
}

// Debug writes debug message with log.Printf
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.printf(DEBUG, format, v...)
}

// Debug writes debug message with log.Println
func (l *Logger) Debugln(v ...interface{}) {
	l.println(DEBUG, v...)
}

// Info writes info message with log.Print
func (l *Logger) Info(v ...interface{}) {
	l.print(INFO, v...)
}

// Info writes info message with log.Printf
func (l *Logger) Infof(format string, v ...interface{}) {
	l.printf(INFO, format, v...)
}

// Info writes info message with log.Println
func (l *Logger) Infoln(v ...interface{}) {
	l.println(INFO, v...)
}

// Warning writes warning message with log.Print
func (l *Logger) Warning(v ...interface{}) {
	l.print(WARNING, v...)
}

// Warning writes warning message with log.Printf
func (l *Logger) Warningf(format string, v ...interface{}) {
	l.printf(WARNING, format, v...)
}

// Warning writes warning message with log.Println
func (l *Logger) Warningln(v ...interface{}) {
	l.println(WARNING, v...)
}

// Error writes error message with log.Print
func (l *Logger) Error(v ...interface{}) {
	l.print(ERROR, v...)
}

// Error writes error message with log.Printf
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.printf(ERROR, format, v...)
}

// Error writes error message with log.Println
func (l *Logger) Errorln(v ...interface{}) {
	l.println(ERROR, v...)
}

// Fatal writes fatal message with log.Fatal
func (l *Logger) Fatal(v ...interface{}) {
	l.print(FATAL, v...)
}

// Fatal writes fatal message with log.Fatalf
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.printf(FATAL, format, v...)
}

// Fatal writes fatal message with log.Fatalln
func (l *Logger) Fatalln(v ...interface{}) {
	l.println(FATAL, v...)
}

// Panic writes panic message with log.Panic
func (l *Logger) Panic(v ...interface{}) {
	l.print(PANIC, v...)
}

// Panic writes panic message with log.Panicf
func (l *Logger) Panicf(format string, v ...interface{}) {
	l.printf(PANIC, format, v...)
}

// Panic writes panic message with log.Panicln
func (l *Logger) Panicln(v ...interface{}) {
	l.println(PANIC, v...)
}

// Debug writes a debug message with the standard logger
func Debug(v ...interface{}) {
	std.Debug(v...)
}

// Debugf writes a debug message with the standard logger
func Debugf(format string, v ...interface{}) {
	std.Debugf(format, v...)
}

// Debugln writes a debug message with the standard logger
func Debugln(v ...interface{}) {
	std.Debugln(v...)
}

// Info writes a info message with the standard logger
func Info(v ...interface{}) {
	std.Info(v...)
}

// Infof writes a info message with the standard logger
func Infof(format string, v ...interface{}) {
	std.Infof(format, v...)
}

// Infoln writes a info message with the standard logger
func Infoln(v ...interface{}) {
	std.Infoln(v...)
}

// Warning writes a warning message with the standard logger
func Warning(v ...interface{}) {
	std.Warning(v...)
}

// Warningf writes a warning message with the standard logger
func Warningf(format string, v ...interface{}) {
	std.Warningf(format, v...)
}

// Warningln writes a warning message with the standard logger
func Warningln(v ...interface{}) {
	std.Warningln(v...)
}

// Error writes a error message with the standard logger
func Error(v ...interface{}) {
	std.Error(v...)
}

// Errorf writes a error message with the standard logger
func Errorf(format string, v ...interface{}) {
	std.Errorf(format, v...)
}

// Errorln writes a error message with the standard logger
func Errorln(v ...interface{}) {
	std.Errorln(v...)
}

// Panic writes a panic message with the standard logger
func Panic(v ...interface{}) {
	std.Panic(v...)
}

// Panicf writes a panic message with the standard logger
func Panicf(format string, v ...interface{}) {
	std.Panicf(format, v...)
}

// Panicln writes a panic message with the standard logger
func Panicln(v ...interface{}) {
	std.Panicln(v...)
}
