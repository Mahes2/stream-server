package http

import (
	"context"
	"net"
	"net/http"

	hasherService "github.com/codespade/stream-server/service/hasher"
)

// Server struct
type Server struct {
	server        *http.Server
	HasherService hasherService.Service
}

// Serve will create, bind, and run a GRPC server
func (s *Server) Serve(port string) error {
	s.server = &http.Server{
		Handler: handler(),
	}

	svc = s.HasherService

	// Create port listener
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	// Start gRPC server
	return s.server.Serve(lis)
}

// GracefulStop gracefully stop server
func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
