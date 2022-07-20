package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	httpapi "github.com/codespade/stream-server/api/http"
)

// runGRPCServer will run GRPC server with specified parameter
func runHTTPServer(httpsrv httpapi.Server, port string) error {
	idleConnsClosed := make(chan struct{})
	go func() {
		signals := make(chan os.Signal, 1)

		// When using socketmaster, it send SIGTERM after spawning new process,
		// SIGHUP is for handling upstart reload
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
		<-signals

		// We received an os signal, shut down.
		if err := httpsrv.Shutdown(context.Background()); err != nil {
			log.Println("HTTP server shutdown: ", err)
		}
		close(idleConnsClosed)
	}()

	log.Println("HTTP server running on port", port)
	if err := httpsrv.Serve(port); err != http.ErrServerClosed {
		// Error starting or closing listener:
		return err
	}

	<-idleConnsClosed
	log.Println("HTTP server stopping")
	return nil
}
