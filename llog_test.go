package llog

import (
	"bufio"
	"bytes"
	"testing"
)

func TestDebug(t *testing.T) {
	Debug("Hallo", "Marc")
}

func TestInfo(t *testing.T) {
	var b bytes.Buffer
	SetWriter(bufio.NewWriter(&b))
	Info("Hallo", "Marc")
	Infof("Hallo %s", "Marc")
	Infoln("Hallo", "Marc")
}
