package dockerutil

import (
	"context"
	"net/http"
	"sync"
)

// StopServers will
func StopServers(ctx context.Context, servers ...*http.Server) chan error {
	errChan := make(chan error, len(servers))
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(servers))
		for _, server := range servers {
			go func(server *http.Server) {
				errChan <- server.Shutdown(ctx)
				wg.Done()
			}(server)
		}
		wg.Wait()
		close(errChan)
	}()

	return errChan
}
