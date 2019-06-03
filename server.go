package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/timberio/tcp_server"
)

var summaryPath = "/tmp/tcp_test_server_summary.json"

type Server struct {
	server          *tcp_server.Server
	ConnectionCount int64  `json:"connection_count"`
	FirstMessage    string `json:"first_message"`
	LastMessage     string `json:"last_message"`
	MessageCount    int64  `json:"message_count"`
	sampleCadence   float64
	sampleMessage   string
}

func (s *Server) Listen() {
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	go func() {
		sig := <-gracefulStop
		log.Printf("Caught sig: %+v", sig)
		s.WriteSummary()
		log.Println("Server stopped")
		os.Exit(0)
	}()

	// Print debug output on an interval. This helps with providing insight
	// into activity without saturating IO.
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				log.Printf("Received %v messages across %v connections", s.MessageCount, s.ConnectionCount)

				if s.sampleMessage != "" {
					log.Printf("Sample: %s", s.sampleMessage)
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	s.server.Listen()
}

func (s *Server) WriteSummary() {
	sBytes, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(summaryPath, sBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Wrote activity summary to %s", summaryPath)
}

func NewServer(address string) *Server {
	os.Remove(summaryPath)

	internal_server := tcp_server.New(address)

	server := &Server{server: internal_server, ConnectionCount: 0, MessageCount: 0, sampleCadence: 5000.0}

	internal_server.OnNewClient(func(c *tcp_server.Client) {
		log.Print("New connection established")
		server.ConnectionCount++
	})

	internal_server.OnNewMessage(func(c *tcp_server.Client, message string) {
		server.MessageCount++

		if server.MessageCount == 1 {
			server.FirstMessage = message
		}

		server.LastMessage = message

		if math.Mod(float64(server.MessageCount), server.sampleCadence) == 0 {
			server.sampleMessage = message
		}
	})

	internal_server.OnClientConnectionClosed(func(c *tcp_server.Client, err error) {
		log.Print("Connection lost")
	})

	return server
}
