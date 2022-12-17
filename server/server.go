package server

import (
	"fmt"
	"github.com/codespade/stream-server/api/grpc"
	"github.com/codespade/stream-server/api/http"
	"github.com/codespade/stream-server/repository"
	"github.com/codespade/stream-server/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //Postgres driver
)

// InitGRPC to initiate all DI for grpc service handler implementation
func InitGRPC(port string) error {
	db, err := InitDB()
	if err != nil {
		fmt.Println(err)
	}

	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	//repository.Db = db
	grpcServer := grpc.Server{
		Service: svc,
	}

	return runGRPCServer(grpcServer, port)
}

func InitHttp(port string) error {
	db, err := InitDB()
	if err != nil {
		fmt.Println(err)
	}

	repo := repository.NewRepository(db)
	svc := service.NewService(repo)

	//repository.Db = db
	httpServer := http.Server{
		Service: svc,
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
