package docker

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

const imagePrefix = "kex"

type Service struct {
	Dir         string
	id          string
	BackendPort int
}

func (s Service) Name() string {
	return filepath.Base(s.Dir)
}

func (s Service) Image() string {
	return fmt.Sprintf("%s-%s", imagePrefix, s.Name())
}

func BuildService(s Service) error {
	cmd := exec.Command("docker", "build", s.Dir, "-t", s.Image())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func (s *Service) Start(port int) error {
	var err error
	s.id, err = start(s.Image(), port)
	return err
}

func start(image string, port int) (string, error) {
	// docker run --add-host=host.docker.internal:host-gateway -p port:80 --rm -d image
	cmd := exec.Command("docker", "run", "--add-host=host.docker.internal:host-gateway", "-p", strconv.Itoa(port)+":80", "--rm", "-d", image)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%s\n", output)
		return "", err
	}

	// Remove newline of ID
	if n := len(output); n > 0 {
		output = output[:n-1]
	}
	return string(output), nil
}

func (s Service) Stop() error {
	return stop(s.id)
}

func stop(id string) error {
	cmd := exec.Command("docker", "stop", "-t", "1", id)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Got error: %s\n", output)
	}
	return err
}
