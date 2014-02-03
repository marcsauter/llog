package llog

import (
	"bufio"
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"testing"
)

const (
	StdHeader  = `\d{4}/[01]\d/[0-3]\d [0-2]\d:[0-5]\d:[0-5]\d\s`
	TestPrefix = `INFO - `
)

var message []interface{} = []interface{}{"Test", "Message", 1}
var format string = "%s %s %d\n"

func TestThreshold(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	logger := New(INFO, w)
	logger.Debug("Hallo", "Marc")
	w.Flush()
	if b.Len() > 0 {
		t.Errorf("threshold does not work\n")
	}
}

func TestPrint(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	logger := New(INFO, w)
	logger.Info(message...)
	w.Flush()
	r := regexp.MustCompile(strings.Join([]string{StdHeader, fmt.Sprint(message...)}, ""))
	if !r.Match(b.Bytes()) {
		t.Errorf("TestPrint: message does not match\n")
	}
}

func TestPrintf(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	logger := New(INFO, w)
	logger.Infof(format, message...)
	w.Flush()
	r := regexp.MustCompile(strings.Join([]string{StdHeader, fmt.Sprintf(format, message...)}, ""))
	if !r.Match(b.Bytes()) {
		t.Errorf("TestPrintf: message does not match\n")
	}
}

func TestPrintln(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	logger := New(INFO, w)
	logger.Infoln(message...)
	w.Flush()
	r := regexp.MustCompile(strings.Join([]string{StdHeader, fmt.Sprintln(message...)}, ""))
	if !r.Match(b.Bytes()) {
		t.Errorf("TestPrintln: message does not match\n")
	}
}

func TestPrefixPrint(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	logger := New(INFO, w)
	logger.Pseverity(true)
	logger.Info(message...)
	w.Flush()
	r := regexp.MustCompile(strings.Join([]string{StdHeader, TestPrefix, fmt.Sprint(message...)}, ""))
	if !r.Match(b.Bytes()) {
		t.Errorf("TestPrefixPrint: message does not match\n")
	}
}

func TestPrefixPrintf(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	logger := New(INFO, w)
	logger.Pseverity(true)
	logger.Infof(format, message...)
	w.Flush()
	r := regexp.MustCompile(strings.Join([]string{StdHeader, TestPrefix, fmt.Sprintf(format, message...)}, ""))
	if !r.Match(b.Bytes()) {
		t.Errorf("TestPrefixPrintf: message does not match\n")
	}
}

func TestPrefixPrintln(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	logger := New(INFO, w)
	logger.Pseverity(true)
	logger.Infoln(message...)
	w.Flush()
	r := regexp.MustCompile(strings.Join([]string{StdHeader, TestPrefix, fmt.Sprintln(message...)}, ""))
	if !r.Match(b.Bytes()) {
		t.Errorf("TestPrefixPrintln: message does not match\n")
	}
}
