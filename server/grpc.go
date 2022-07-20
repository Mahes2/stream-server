package server

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	grpcapi "github.com/codespade/stream-server/api/grpc"
)

// runGRPCServer will run GRPC server with specified parameter
func runGRPCServer(grpcsrv grpcapi.Server, port string) error {
	idleConnsClosed := make(chan struct{})
	go func() {
		signals := make(chan os.Signal, 1)

		// When using socketmaster, it send SIGTERM after spawning new process,
		// SIGHUP is for handling upstart reload
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
		<-signals

		// We received an os signal, shut down.
		grpcsrv.GracefulStop()
		log.Println("GRPC server shutdown gracefully")
		close(idleConnsClosed)
	}()

	log.Println("GRPC server running on port", port)
	if err := grpcsrv.Serve(port); err != http.ErrServerClosed {
		// Error starting or closing listener:
		return err
	}

	<-idleConnsClosed
	log.Println("GRPC server stopping")
	return nil
}
