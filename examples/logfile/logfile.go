// logfile is example for logging to a logfile
package main

import (
	"github.com/marcsauter/llog"
	"io/ioutil"
)

func main() {
	filename := "example.log"
	w, _ := llog.AddLogfile(ioutil.Discard, filename)
	logger := llog.New(llog.ERROR, w)
	// Info should not appear in example.log
	logger.Info("Info message")
	// Error should be written to example.log
	logger.Error("Error message")

}
