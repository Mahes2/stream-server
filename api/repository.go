package api

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	BlockID(ctx context.Context, id string) (Response, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	r := repository{
		db: db,
	}

	return &r
}

const Query = `
UPDATE test
SET status = $1
WHERE id = $2
RETURNING id, status
`

func (r *repository) BlockID(ctx context.Context, id string) (Response, error) {
	row := r.db.DB.QueryRowContext(ctx, Query, "BLOCKED", id)

	var (
		i Response
	)

	err := row.Scan(
		&i.Id,
		&i.Status,
	)

	return i, err
}
