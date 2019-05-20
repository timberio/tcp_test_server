package main

import (
	"log"
	"os"

	"github.com/timberio/tcp_server"
)

type Server struct {
	server *tcp_server.Server
	File   *os.File
}

func (s *Server) Listen() {
	s.server.Listen()
}

func NewServer(address string, filePath string) *Server {
	internal_server := tcp_server.New(address)

	server := &Server{server: internal_server}

	if filePath != "" {
		f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		server.File = f
	}

	internal_server.OnNewClient(func(c *tcp_server.Client) {
		log.Print("New connection established")
	})

	internal_server.OnNewMessage(func(c *tcp_server.Client, message string) {
		if server.File != nil {
			_, err := server.File.WriteString(message)
			if err != nil {
				log.Fatal(err)
			}
		}
	})

	internal_server.OnClientConnectionClosed(func(c *tcp_server.Client, err error) {
		log.Print("Connection lost")
	})

	return server
}
