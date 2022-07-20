package http

import (
	"context"
	"net"
	"net/http"

	"github.com/codespade/stream-server/api"
	controller "github.com/codespade/stream-server/api/http/controller"
)

// Server struct
type Server struct {
	server     *http.Server
	Repository api.Repository
}

// Serve will create, bind, and run a GRPC server
func (s *Server) Serve(port string) error {
	controller.Init(s.Repository)

	s.server = &http.Server{
		Handler: handler(),
	}

	// Create port listener
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	// Start gRPC server
	return s.server.Serve(lis)
}

//GracefulStop gracefully stop server
func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
