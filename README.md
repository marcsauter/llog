llog
----
Level logging on top of the standard log package. 

A simple example:

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

For more examples see example folder.

Documentation
-------------
[http://godoc.org/github.com/marcsauter/llog][1]


  [1]: http://godoc.org/github.com/marcsauter/llog