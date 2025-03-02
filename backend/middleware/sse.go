package middleware

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type SSEMsg struct {
	Name string `json:"name"`
	Data any    `json:"data"`
}

type SSE struct {
	Message chan SSEMsg

	conns     map[chan string]bool
	NewConn   chan chan string
	CloseConn chan chan string
}

// NewSSE creates a new SSE instance and starts listening for new connections
func NewSSE() *SSE {
	sse := &SSE{
		Message:   make(chan SSEMsg),
		conns:     make(map[chan string]bool),
		NewConn:   make(chan chan string),
		CloseConn: make(chan chan string),
	}
	go sse.Listen()
	return sse
}

func (s *SSE) Listen() {
	for {
		select {
		case conn := <-s.NewConn:
			s.conns[conn] = true
		case conn := <-s.CloseConn:
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
		s.NewConn <- conn

		defer func() {
			s.CloseConn <- conn
		}()

		for msg := range conn {
			c.Writer.Write([]byte(msg))
			c.Writer.Flush()
		}
	}
}
