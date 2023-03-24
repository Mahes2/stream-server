package repository

import (
	"context"
	"github.com/codespade/stream-server/entity"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	BlockID(ctx context.Context, id string) (entity.Response, error)
}

type repository struct {
	Db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{
		Db: db,
	}
}

const Query = `
UPDATE test
SET status = $1
WHERE id = $2
RETURNING id, status
`

func (r *repository) BlockID(ctx context.Context, id string) (entity.Response, error) {
	row := r.Db.QueryRowContext(ctx, Query, "BLOCKED", id)

	var i entity.Response

	err := row.Scan(
		&i.Id,
		&i.Status,
	)

	return i, err
}
