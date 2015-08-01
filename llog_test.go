package llog

import (
	"bufio"
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

const (
	StdHeader     = `\d{4}/[01]\d/[0-3]\d [0-2]\d:[0-5]\d:[0-5]\d\s`
	DebugPrefix   = `DEBUG - `
	InfoPrefix    = `INFO - `
	WarningPrefix = `WARNING - `
	ErrorPrefix   = `ERROR - `
	FatalPrefix   = `FATAL - `
	PanicPrefix   = `PANIC - `
)

var message []interface{} = []interface{}{"Test", "Message", 1}
var printmsg string = fmt.Sprint(message...)
var printfmsg string = fmt.Sprintf("%s %s %d\n", message...)
var printlnmsg string = fmt.Sprintln(message...)

func print(s severity, m string, ps bool) *bytes.Buffer {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	l := New(s, w)
	l.Pseverity(ps)
	reflect.ValueOf(l).MethodByName(m).Call([]reflect.Value{reflect.ValueOf(printmsg)})
	w.Flush()
	return &b
}

func printf(s severity, m string, ps bool) *bytes.Buffer {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	l := New(s, w)
	l.Pseverity(ps)
	reflect.ValueOf(l).MethodByName(m).Call([]reflect.Value{reflect.ValueOf(printfmsg)})
	w.Flush()
	return &b
}

func println(s severity, m string, ps bool) *bytes.Buffer {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	l := New(s, w)
	l.Pseverity(ps)
	reflect.ValueOf(l).MethodByName(m).Call([]reflect.Value{reflect.ValueOf(printlnmsg)})
	w.Flush()
	return &b
}

func TestThreshold(t *testing.T) {
	if print(INFO, "Debug", false).Len() > 0 {
		t.Errorf("threshold does not work\n")
	}
}

func TestPrint(t *testing.T) {
	r := regexp.MustCompile(strings.Join([]string{StdHeader, printmsg}, ""))
	m := print(INFO, "Info", false)
	fmt.Print(m.String())
	if !r.Match(m.Bytes()) {
		t.Errorf("message does not match: %s\n", m.String())
	}
}

func TestPrintf(t *testing.T) {
	r := regexp.MustCompile(strings.Join([]string{StdHeader, printfmsg}, ""))
	m := printf(INFO, "Infof", false)
	if !r.Match(m.Bytes()) {
		t.Errorf("message does not match: %s\n", m.String())
	}
}

func TestPrintln(t *testing.T) {
	r := regexp.MustCompile(strings.Join([]string{StdHeader, printlnmsg}, ""))
	m := println(INFO, "Infoln", false)
	if !r.Match(m.Bytes()) {
		t.Errorf("message does not match: %s\n", m.String())
	}
}

func TestDebugPrint(t *testing.T) {
	r := regexp.MustCompile(strings.Join([]string{StdHeader, DebugPrefix, printmsg}, ""))
	m := print(DEBUG, "Debug", true)
	if !r.Match(m.Bytes()) {
		t.Errorf("message does not match: %s\n", m.String())
	}
}

func TestDebugPrintf(t *testing.T) {
	r := regexp.MustCompile(strings.Join([]string{StdHeader, DebugPrefix, printfmsg}, ""))
	m := printf(DEBUG, "Debugf", true)
	if !r.Match(m.Bytes()) {
		t.Errorf("message does not match: %s\n", m.String())
	}
}

func TestDebugPrintln(t *testing.T) {
	r := regexp.MustCompile(strings.Join([]string{StdHeader, DebugPrefix, printlnmsg}, ""))
	m := println(DEBUG, "Debugln", true)
	if !r.Match(m.Bytes()) {
		t.Errorf("message does not match: %s\n", m.String())
	}
}
