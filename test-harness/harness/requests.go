package harness

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

const requestDir = "requests"

type Request struct {
	Name string
	Data []byte
}

func LoadRequests(files []string) ([]Request, error) {
	var reqs []Request
	for _, f := range files {
		if !strings.HasSuffix(f, ".req") {
			continue
		}
		b, err := ioutil.ReadFile(f)
		if err != nil {
			return nil, err
		}
		reqs = append(reqs, Request{
			Name: strings.TrimSuffix(filepath.Base(f), ".req"),
			Data: b,
		})
	}
	sort.Slice(reqs, func(i, j int) bool {
		return reqs[i].Name < reqs[j].Name
	})
	return reqs, nil
}

func sendRequest(req Request, port int) ([]byte, error) {
	conn, err := net.Dial("tcp", "localhost:"+strconv.Itoa(port))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	n, err := conn.Write(req.Data)
	if err != nil {
		return nil, fmt.Errorf("error while writing req: %v", err)
	}
	if n != len(req.Data) {
		return nil, errors.New("failed to write entire request")
	}

	conn.SetReadDeadline(time.Now().Add(2000 * time.Millisecond))
	resp, err := ioutil.ReadAll(conn)
	if err != nil && strings.Index(err.Error(), "connection reset by peer") != -1 {
		fmt.Printf("error while reading resp, continuing: %v\n", err)
		return resp, nil
	}

	if err != nil && !errors.Is(err, os.ErrDeadlineExceeded) {
		return nil, fmt.Errorf("error while reading resp: %v", err)
	}
	return resp, nil
}
