package dockerutil

import (
	"testing"
	"os"
	"syscall"
	"time"
)

func TestHandleDockerStop(t *testing.T) {
	stopChan := make(chan struct{})
	sigChan := HandleDockerStop(func() {
		stopChan<-struct{}{}
	})
	sigChan <- syscall.SIGTERM
	select {
	case <-stopChan:
	case <-time.After(2 * time.Second):
		t.Fail()
	}
}

func TestHandleSignalChannel(t *testing.T) {
	sigChan := make(chan os.Signal, 1)
	stopChan := make(chan struct{})
	HandleSignalChannel(sigChan, func() {
		stopChan<-struct{}{}
	})
	sigChan <- syscall.SIGTERM
	select {
		case <-stopChan:
		case <-time.After(2 * time.Second):
			t.Fail()
	}
}