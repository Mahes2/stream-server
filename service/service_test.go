package service

import (
	"context"
	"github.com/codespade/stream-server/entity"
	"github.com/codespade/stream-server/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var repo = repository.Mock{Mock: mock.Mock{}}
var svc = NewService(&repo)

func TestService_BlockID(t *testing.T) {
	dummy := entity.Response{
		Id:     "5a5fa251",
		Status: "BLOCKED",
	}
	repo.Mock.On("BlockID", "5a5fa251").Return(dummy)

	resp, err := svc.BlockID(context.Background(), "5a5fa251")
	assert.Equal(t, dummy, resp)
	assert.Nil(t, err)
}

func TestService_HashToMD5(t *testing.T) {
	resp := svc.HashToMD5("123")
	assert.Equal(t, "202cb962ac59075b964b07152d234b70", resp)
}

func BenchmarkService_BlockID(b *testing.B) {
	b.Run("TestService_BlockID", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := svc.BlockID(context.Background(), "5a5fa251")
			if err != nil {
				return
			}
		}
	})
}

func BenchmarkService_HashToMD5(b *testing.B) {
	b.Run("TestService_HashToMD5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			svc.HashToMD5("123")
		}
	})
}
