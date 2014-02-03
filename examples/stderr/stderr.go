package main

import (
	"github.com/marcsauter/llog"
	"os"
)

func main() {
	logger := llog.New(llog.ERROR, os.Stderr)
	logger.Info("Info message")
	logger.Error("Error message")
	logger.SetThreshold(llog.INFO)
	logger.Info("Another info message")
}
