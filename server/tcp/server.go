package tcp

import (
	"net"
	"whisky-server/server"
)

type Server struct {
	handle map[uint]func(in interface{}) (interface{}, error)
}

func NewServer() server.Server {
	return &Server{
		handle: make(map[uint]func(in interface{}) (interface{}, error)),
	}
}

func (s Server) Server() error {
	l, err := net.Listen("tcp",":3207")
	if err != nil {
		return err
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		go s.handleConnection(conn)
	}
}

func (s Server) handleConnection(conn net.Conn) {
	tempBuffer := make([]byte, 0)
	readerChannel := make(chan []byte, 16)

	buffer := make([]byte, 1024)
	var count int32
	for {
		count++
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}
		tempBuffer = Unpack(append(tempBuffer, buffer[:n]...), readerChannel)
	}
}

func (s Server) Stop() error {
	return nil
}
