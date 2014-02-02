package llog

import (
	"fmt"
	"io"
	"log"
	"log/syslog"
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

type loggerT struct {
	threshold severity
	prefix    *string
	writer    io.Writer
	syslog    *syslog.Writer
}

func (l *loggerT) print(s severity, v ...interface{}) {
	if s < l.threshold {
		return
	}
	log.Print(v...)
	l.syslogMessage(s, fmt.Sprint(v...))
	l.action(s, v...)
}

func (l *loggerT) printf(s severity, format string, v ...interface{}) {
	if s < l.threshold {
		return
	}
	log.Printf(format, v...)
	l.syslogMessage(s, fmt.Sprintf(format, v...))
	l.action(s, v...)
}

func (l *loggerT) println(s severity, v ...interface{}) {
	if s < l.threshold {
		return
	}
	log.Println(v...)
	l.syslogMessage(s, fmt.Sprintln(v...))
	l.action(s, v...)
}

func (l *loggerT) syslogMessage(s severity, m string) {
	if logger.syslog != nil {
		switch s {
		case DEBUG:
			logger.syslog.Debug(m)
		case INFO:
			logger.syslog.Info(m)
		case WARNING:
			logger.syslog.Warning(m)
		case ERROR:
			logger.syslog.Err(m)
		case FATAL:
			logger.syslog.Alert(m)
		case PANIC:
			logger.syslog.Emerg(m)
		}
	}
}

func (l *loggerT) action(s severity, v ...interface{}) {
	switch s {
	case ERROR:
		os.Exit(1)
	case PANIC:
		s := fmt.Sprint(v...)
		panic(s)
	}
}

func Threshold() severity {
	return logger.threshold
}

func SetThreshold(s severity) {
	logger.threshold = s
}

func Hdate(set bool) {
	if set {
		log.SetFlags(log.Flags() | log.Ldate)
	} else {
		log.SetFlags(log.Flags() & ^log.Ldate)
	}
}

func Htime(set bool) {
	if set {
		log.SetFlags(log.Flags() | log.Ltime)
	} else {
		log.SetFlags(log.Flags() & ^log.Ltime)
	}
}

func Hmicroseconds(set bool) {
	if set {
		log.SetFlags(log.Flags() | log.Lmicroseconds)
	} else {
		log.SetFlags(log.Flags() & ^log.Lmicroseconds)
	}
}

func Hshortfile(set bool) {
	if set {
		log.SetFlags(log.Flags() & ^log.Llongfile)
		log.SetFlags(log.Flags() | log.Lshortfile)
	} else {
		log.SetFlags(log.Flags() & ^log.Lshortfile)
	}
}

func Hlongfile(set bool) {
	if set {
		log.SetFlags(log.Flags() & ^log.Lshortfile)
		log.SetFlags(log.Flags() | log.Llongfile)
	} else {
		log.SetFlags(log.Flags() & ^log.Llongfile)
	}
}

func Hseverity(set bool) {

}

func SetOutput(w io.Writer) {
	logger.writer = w
	log.SetOutput(logger.writer)
}

func SetLogfile(filename string) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	logger.writer = io.MultiWriter(logger.writer, f)
	log.SetOutput(logger.writer)
	return nil
}

func SetSyslog(priority syslog.Priority, tag string) error {
	// set default priority
	if priority == 0 {
		priority = syslog.LOG_INFO | syslog.LOG_USER
	}
	w, err := syslog.New(priority, tag)
	if err != nil {
		return err
	}
	logger.syslog = w
	return nil
}

var logger *loggerT

func init() {
	logger = new(loggerT)
	// initial log level
	logger.threshold = INFO
	logger.writer = io.MultiWriter(os.Stderr)
	log.SetOutput(logger.writer)
}

func Debug(v ...interface{}) {
	logger.print(DEBUG, v...)
}

func Debugf(format string, v ...interface{}) {
	logger.printf(DEBUG, format, v...)
}

func Debugln(v ...interface{}) {
	logger.println(DEBUG, v...)
}

func Info(v ...interface{}) {
	logger.print(INFO, v...)
}

func Infof(format string, v ...interface{}) {
	logger.printf(INFO, format, v...)
}

func Infoln(v ...interface{}) {
	logger.println(INFO, v...)
}

func Warning(v ...interface{}) {
	logger.print(WARNING, v...)
}

func Warningf(format string, v ...interface{}) {
	logger.printf(WARNING, format, v...)
}

func Warningln(v ...interface{}) {
	logger.println(WARNING, v...)
}

func Error(v ...interface{}) {
	logger.print(ERROR, v...)
}

func Errorf(format string, v ...interface{}) {
	logger.printf(ERROR, format, v...)
}

func Errorln(v ...interface{}) {
	logger.println(ERROR, v...)
}

func Fatal(v ...interface{}) {
	logger.print(FATAL, v...)
}

func Fatalf(format string, v ...interface{}) {
	logger.printf(FATAL, format, v...)
}

func Fatalln(v ...interface{}) {
	logger.println(FATAL, v...)
}

func Panic(v ...interface{}) {
	logger.print(PANIC, v...)
}

func Panicf(format string, v ...interface{}) {
	logger.printf(PANIC, format, v...)
}

func Panicln(v ...interface{}) {
	logger.println(PANIC, v...)
}
