package server

import (
	"fmt"

	"github.com/codespade/stream-server/api/grpc"
	"github.com/codespade/stream-server/api/http"
	orderRepo "github.com/codespade/stream-server/repository/order"
	testRepo "github.com/codespade/stream-server/repository/test"
	hasherService "github.com/codespade/stream-server/service/hasher"
	orderService "github.com/codespade/stream-server/service/order"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //Postgres driver
)

// InitGRPC to initiate all DI for grpc service handler implementation
func InitGRPC(port string) error {
	db, err := InitDB()
	if err != nil {
		fmt.Println(err)
	}

	testRepo := testRepo.NewRepository(db)
	hasherSvc := hasherService.NewService(testRepo)

	orderRepo := orderRepo.NewRepository(db)
	orderSvc := orderService.NewService(orderRepo)

	grpcServer := grpc.Server{
		HasherService: hasherSvc,
		OrderService:  orderSvc,
	}

	return runGRPCServer(grpcServer, port)
}

func InitHttp(port string) error {
	db, err := InitDB()
	if err != nil {
		fmt.Println(err)
	}

	repo := testRepo.NewRepository(db)
	svc := hasherService.NewService(repo)

	//repository.Db = db
	httpServer := http.Server{
		HasherService: svc,
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
