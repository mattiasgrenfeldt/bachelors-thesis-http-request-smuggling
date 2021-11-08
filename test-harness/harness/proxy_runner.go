package harness

import (
	"fmt"
	"time"

	"github.com/mattiasgrenfeldt/bachelors-thesis-http-request-smuggling/test-harness/harness/backend"
	"github.com/mattiasgrenfeldt/bachelors-thesis-http-request-smuggling/test-harness/harness/docker"
)

func runProxy(proxy docker.Service, port int, reqs []Request) ([]Response, error) {
	var resps []Response
	for i, r := range reqs {
		fmt.Printf("[%12s] req %d/%d %s\n", proxy.Name(), i+1, len(reqs), r.Name)
		b := backend.New()
		b.Start(proxy.BackendPort)
		if err := proxy.Start(port); err != nil {
			return nil, err
		}
		time.Sleep(2000 * time.Millisecond)
		d, err := sendRequest(r, port)
		if err != nil {
			return nil, err
		}
		if err := proxy.Stop(); err != nil {
			return nil, fmt.Errorf("error while shutting down proxy: %v", err)
		}
		backendData, err := b.Stop()
		if err != nil {
			return nil, fmt.Errorf("error while shutting down backend: %v", err)
		}
		resps = append(resps, Response{
			Name:        r.Name,
			Data:        d,
			BackendData: backendData,
		})
	}
	return resps, nil
}
