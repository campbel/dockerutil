package dockerutil

import (
	"os"
	"os/signal"
	"syscall"
)

// HandleDockerStop will run the passed function on a SIGTERM signal
func HandleDockerStop(handler func()) chan os.Signal {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM)
	HandleSignalChannel(sigChan, handler)
	return sigChan
}

// HandleSignalChannel will run the passed function when the signal channel fires
func HandleSignalChannel(sigChan chan os.Signal, handler func()) {
	go func() {
		<-sigChan
		handler()
	}()
}
