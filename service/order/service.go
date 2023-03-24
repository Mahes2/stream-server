package order

import (
	"context"

	"github.com/codespade/stream-server/entity"
	orderRepo "github.com/codespade/stream-server/repository/order"
)

type Service interface {
	ListOrder(ctx context.Context, driverId int64) ([]entity.Order, error)
}

type service struct {
	repo orderRepo.Repository
}

func NewService(r orderRepo.Repository) *service {
	return &service{
		repo: r,
	}
}

func (s *service) ListOrder(ctx context.Context, driverId int64) ([]entity.Order, error) {
	return s.repo.ListOrder(ctx, driverId)
}
