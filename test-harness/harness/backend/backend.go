package backend

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strconv"
	"strings"
)

type Backend struct {
	listener net.Listener
	barrier  chan int
	data     []byte
	err      error
}

func New() Backend {
	return Backend{
		barrier: make(chan int),
	}
}

func (b *Backend) Start(port int) error {
	l, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(port))
	if err != nil {
		return err
	}
	b.listener = l
	go b.listenAndHandle(l)
	return nil
}

var response = []byte("HTTP/1.0 200 OK\r\nContent-Type: text/html\r\nConnection: close\r\nContent-Length: 18\r\n\r\nhello from backend")

func (b *Backend) listenAndHandle(l net.Listener) {
	defer func() {
		l.Close()
		b.barrier <- 0
	}()

	// Let's assume we only get one connection.
	// We have to be careful that this doesn't cause a problem.
	conn, err := l.Accept()
	if err != nil {
		// Let's ignore these errors so that we can close from the outside
		// b.err = err
		return
	}
	defer conn.Close()

	var n int
	n, b.err = conn.Write(response)
	if b.err != nil {
		return
	}
	if n != len(response) {
		b.err = fmt.Errorf("didn't write all bytes, wrote: %v payload: %v", n, len(response))
	}

	b.data, b.err = ioutil.ReadAll(conn)
}

func (b *Backend) Stop() ([]byte, error) {
	if b.listener != nil {
		b.listener.Close()
	}
	<-b.barrier
	if b.err != nil && strings.Index(b.err.Error(), "connection reset by peer") != -1 {
		log.Printf("Backend got error, continuing anyways: %v", b.err)
		return b.data, nil
	}
	if b.err != nil {
		return nil, b.err
	}
	return b.data, nil
}
