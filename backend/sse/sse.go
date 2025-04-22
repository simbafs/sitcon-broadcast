package sse

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type Msg struct {
	Topic []string `json:"topic"`
	Data  any      `json:"data"`
}

type SSE struct {
	Message chan Msg

	conns     map[chan string]bool
	newConn   chan chan string
	closeConn chan chan string
}

// NewSSE creates a new SSE instance and starts listening for new connections
func New() *SSE {
	sse := &SSE{
		Message:   make(chan Msg),
		conns:     make(map[chan string]bool),
		newConn:   make(chan chan string),
		closeConn: make(chan chan string),
	}
	go sse.Listen()
	return sse
}

func (s *SSE) Listen() {
	for {
		select {
		case conn := <-s.newConn:
			s.conns[conn] = true
		case conn := <-s.closeConn:
			delete(s.conns, conn)
		case msg := <-s.Message:
			data, err := encodeSSEMsg(msg)
			if err != nil {
				continue
			}
			for conn := range s.conns {
				conn <- data
			}
		}
	}
}

func encodeSSEMsg(msg any) (string, error) {
	data, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}

	return "data: " + string(data) + "\n\n", nil
}

func (s *SSE) GinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// set headers
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Header().Set("Transfer-Encoding", "chunked")

		// craete connection
		conn := make(chan string)
		s.newConn <- conn

		defer func() {
			s.closeConn <- conn
		}()

		for msg := range conn {
			c.Writer.Write([]byte(msg))
			c.Writer.Flush()
		}
	}
}
