package harness

import (
	"io/ioutil"
	"path/filepath"
	"strconv"

	"github.com/mattiasgrenfeldt/bachelors-thesis-http-request-smuggling/test-harness/harness/docker"
)

const serviceDir = "services"

func BuildAll() (proxies []docker.Service, servers []docker.Service, _ error) {
	proxies, err := buildDir(filepath.Join(serviceDir, "proxies"))
	if err != nil {
		return nil, nil, err
	}
	readBackendPorts(proxies)

	servers, err = buildDir(filepath.Join(serviceDir, "servers"))
	if err != nil {
		return nil, nil, err
	}
	return proxies, servers, nil
}

func buildDir(dir string) ([]docker.Service, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var services []docker.Service
	for _, f := range files {
		if !f.IsDir() {
			continue
		}
		s := docker.Service{
			Dir: filepath.Join(dir, f.Name()),
		}
		if err := docker.BuildService(s); err != nil {
			return nil, err
		}
		services = append(services, s)
	}
	return services, nil
}

func readBackendPorts(proxies []docker.Service) error {
	for i, p := range proxies {
		path := filepath.Join(p.Dir, "BACKEND_PORT")
		b, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		port, err := strconv.Atoi(string(b))
		if err != nil {
			return err
		}
		proxies[i].BackendPort = port
	}
	return nil
}
