package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/codespade/stream-server/entity"
	"github.com/codespade/stream-server/repository"
)

type Service interface {
	BlockID(ctx context.Context, id string) (entity.Response, error)
	HashToMD5(id string) string
}

type service struct {
	repo repository.Repository
}

func NewService(r repository.Repository) *service {
	return &service{
		repo: r,
	}
}

func (s *service) BlockID(ctx context.Context, id string) (entity.Response, error) {
	return s.repo.BlockID(ctx, id)
}

func (s *service) HashToMD5(id string) string {
	hash := md5.New()
	hash.Write([]byte(id))
	data := hex.EncodeToString(hash.Sum(nil))

	return data
}
