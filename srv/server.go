package srv

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strconv"
	"sync"
)

type Server struct {
	mu            sync.Mutex
	ln            []net.Listener
	done          chan struct{}
	concurrencyCh chan struct{}
	open          int32
	MaxConnPerIP  int
	Host          string
	Port          string
	Worker        int
}

type Config struct {
	Host   string
	Port   int
	Worker int
}

func New(config *Config) *Server {

	return &Server{
		Host:   config.Host,
		Port:   strconv.Itoa(config.Port),
		Worker: config.Worker,
	}
}

func (s *Server) ListenAndServe() error {

	if s.Host == "" {
		s.Host = "127.0.0.1"
	}

	if s.Port == "" {
		return errors.New("port required")
	}

	var maxWorker = DEFAULT_WORKER

	if s.Worker != 0 {
		maxWorker = s.Worker
	}

	wp := &Worker{
		WorkerFunc: handleRequest,
		MaxWorker: maxWorker,
		Conn: make(chan net.Conn),
	}
	wp.Start()

	listener, err := net.Listen("tcp4", fmt.Sprintf("%s:%s", s.Host, s.Port))
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		wp.Add(conn)
	}
}

func handleRequest(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			conn.Close()
			return
		}
		fmt.Printf("Message incoming: %s", string(message))
		conn.Write([]byte("Message received.\n"))
	}
}
