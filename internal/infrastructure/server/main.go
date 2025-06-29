package server

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	*http.Server
}

func NewServer(addr string, hdlr http.Handler) *Server {
	return &Server{
		&http.Server{
			Addr:           ":" + addr,
			Handler:        hdlr,
			MaxHeaderBytes: 1 << 20, //1MB
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
	}
}

func (s *Server) Start() error {

	err := s.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start http server: %v", err)
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {

	err := s.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("and error occure while shutdown server: %v", err)
	}

	return nil
}
