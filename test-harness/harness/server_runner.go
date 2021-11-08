package harness

import (
	"fmt"
	"time"

	"github.com/mattiasgrenfeldt/bachelors-thesis-http-request-smuggling/test-harness/harness/docker"
)

func runServer(server docker.Service, port int, reqs []Request) ([]Response, error) {
	var resps []Response
	for i, r := range reqs {
		fmt.Printf("[%12s] req %d/%d %s\n", server.Name(), i+1, len(reqs), r.Name)
		if err := server.Start(port); err != nil {
			return nil, fmt.Errorf("error while starting %v on %d: %v", server.Name(), port, err)
		}
		time.Sleep(500 * time.Millisecond)
		d, err := sendRequest(r, port)
		if err != nil {
			return nil, err
		}
		resps = append(resps, Response{
			Name: r.Name,
			Data: d,
		})
		if err := server.Stop(); err != nil {
			return nil, fmt.Errorf("error while shutting down %v: %v", server.Name(), err)
		}
	}
	return resps, nil
}
