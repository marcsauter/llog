package main

import (
	"github.com/marcsauter/llog"
	"os"
)

func main() {
	filename := "example.log"
	w, _ := llog.AddLogfile(os.Stderr, filename)
	logger := llog.New(llog.INFO, w)
	// Info should not appear neither in example.log nor on stderr
	logger.Debug("Debug message")
	// Error should be written to example.log and stderr
	logger.Error("Error message")
}
