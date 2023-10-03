package main

import (
	"fmt"
	"net"
)

const MSG_SIZE = 512

type SocketServerHandler func(*SocketServer, net.Conn)

type SocketServer struct {
	name, port string
	handler    SocketServerHandler
	in, out    chan []byte
}

func (s *SocketServer) Setup(port string, handler SocketServerHandler) {
	s.name = port
	s.port = port
	s.handler = handler
	s.in = make(chan []byte)
	s.out = make(chan []byte)
}

func (s *SocketServer) Serve() {
	connChan := make(chan net.Conn)

	ln, err := net.Listen("tcp", s.port)
	if err != nil {
		panic(err)
	}

	// connection queue
	go func(connChan chan net.Conn, input chan []byte) {
		var connections [100]net.Conn
		counter := 0

		for {
			select {
			case conn := <-connChan:
				connections[counter] = conn
				counter = counter + 1
			case data := <-input:
				for i := 0; i < counter; i = i + 1 {
					connections[i].Write(data)
				}
			}
		}
	}(connChan, s.in)

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go s.handler(s, conn)
		connChan <- conn
	}
}

func EchoHandler(server *SocketServer, conn net.Conn) {
	recvChan := make(chan []byte)
	errChan := make(chan error)

	// goroutine for receiving from client
	go func(recvChan chan []byte, errChan chan error) {
		for {
			data := make([]byte, MSG_SIZE)

			_, err := conn.Read(data)
			if err != nil {
				errChan <- err
				return
			}

			recvChan <- data
		}
	}(recvChan, errChan)

	// listen to broadcast
	for {
		select {
		// receive from client, write to broadcast
		case data := <-recvChan:
			server.out <- data
		// receive errors, exit
		case <-errChan:
			conn.Close()
			return
		}
	}
}

func main() {
	sa := new(SocketServer)
	sb := new(SocketServer)

	sa.Setup(":1234", EchoHandler)
	sb.Setup(":1235", EchoHandler)

	go sa.Serve()
	go sb.Serve()

	// a broker
	for {
		select {
		case data := <-sa.out:
			fmt.Println("Hear from ", sa.name, string(data))
			sb.in <- data
		case data := <-sb.out:
			fmt.Println("Hear from ", sb.name, string(data))
			sa.in <- data
		}
	}
}
