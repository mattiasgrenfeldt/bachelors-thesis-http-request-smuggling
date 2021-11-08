package harness

import (
	"errors"

	"github.com/mattiasgrenfeldt/bachelors-thesis-http-request-smuggling/test-harness/harness/docker"
)

const portStart = 8000

func RunServices(serviceType string, services []docker.Service, reqs []Request) (map[string][]Response, error) {
	type res struct {
		Name  string
		Resps []Response
		Err   error
	}
	c := make(chan res)
	defer close(c)
	for i, s := range services {
		// Don't remove the two lines below!
		s := s
		i := i
		go func() {
			var r []Response
			var err error
			if serviceType == "server" {
				r, err = runServer(s, portStart+i, reqs)
			} else if serviceType == "proxy" {
				r, err = runProxy(s, portStart+i, reqs)
			} else {
				err = errors.New("bad serviceType")
			}
			c <- res{
				Name:  s.Name(),
				Resps: r,
				Err:   err,
			}
		}()
	}

	result := map[string][]Response{}
	for i := 0; i < len(services); i++ {
		r := <-c
		if r.Err != nil {
			return nil, r.Err
		}
		result[r.Name] = r.Resps
	}
	return result, nil
}
