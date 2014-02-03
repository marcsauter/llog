// syslog is an example for logging to syslog service
// you need to customize priority and level to meet
// the needs of your syslog service
package main

import (
	"github.com/marcsauter/llog"
	"io/ioutil"
	"os"
)

func main() {
	logger := llog.New(llog.INFO, ioutil.Discard)
	// priority 0 defaults to syslog.LOG_INFO | syslog.LOG_USER
	logger.SetSyslog(0, os.Args[0])
	logger.Debug("Debug message")
	logger.Info("Info message")
}
