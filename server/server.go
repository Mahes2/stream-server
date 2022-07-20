package server

import (
	"fmt"

	"github.com/codespade/stream-server/api"
	"github.com/codespade/stream-server/api/grpc"
	"github.com/codespade/stream-server/api/http"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //Postgres driver
)

// InitGRPC to initiate all DI for grpc service handler implementation
func InitGRPC(port string) error {
	db, err := InitDB()
	if err != nil {
		fmt.Println(err)
	}
	repository := api.NewRepository(db)

	grpcServer := grpc.Server{
		Repository: repository,
	}

	return runGRPCServer(grpcServer, port)
}

func InitHttp(port string) error {
	db, err := InitDB()
	if err != nil {
		fmt.Println(err)
	}
	repository := api.NewRepository(db)

	httpServer := http.Server{
		Repository: repository,
	}

	return runHTTPServer(httpServer, port)
}

func InitDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "port=5432 user=postgres password=123456 dbname=postgres sslmode=disable")
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(10)

	return db, nil
}
