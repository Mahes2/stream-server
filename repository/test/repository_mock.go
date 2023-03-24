package repository

import (
	"context"
	"errors"
	"github.com/codespade/stream-server/entity"
	"github.com/stretchr/testify/mock"
)

type Mock struct {
	Mock mock.Mock
}

func (m *Mock) BlockID(_ context.Context, id string) (entity.Response, error) {
	arg := m.Mock.Called(id)
	if arg[0] != nil {
		return arg[0].(entity.Response), nil
	}

	return entity.Response{}, errors.New("empty")
}
