package grpc

import (
	"net"

	hasherService "github.com/codespade/stream-server/service/hasher"
	orderService "github.com/codespade/stream-server/service/order"

	pb "github.com/codespade/stream-server/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server struct
type Server struct {
	pb.UnimplementedStreamServer
	pb.UnimplementedOrderServer
	server *grpc.Server

	HasherService hasherService.Service
	OrderService  orderService.Service
}

// Serve will create, bind, and run a GRPC server
func (s *Server) Serve(port string) error {
	s.server = grpc.NewServer()

	pb.RegisterStreamServer(s.server, s)
	pb.RegisterOrderServer(s.server, s)

	reflection.Register(s.server)

	// Create port listener
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	// Start gRPC server
	return s.server.Serve(lis)
}

// GracefulStop gracefully stop server
func (s *Server) GracefulStop() {
	s.server.GracefulStop()
}
