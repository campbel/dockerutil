package dockerutil

import (
	"context"
	"net/http"
	"testing"
)

func TestStopServer(t *testing.T) {
	errChan := StopServers(context.Background(), &http.Server{}, &http.Server{})
	for err := range errChan {
		if err != nil {
			t.Error(err)
		}
	}
}
