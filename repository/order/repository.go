package order

import (
	"context"

	"github.com/codespade/stream-server/entity"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	ListOrder(ctx context.Context, driverId int64) ([]entity.Order, error)
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
SELECT id, quantity, customer_address FROM "order"
`

func (r *repository) ListOrder(ctx context.Context, driverId int64) ([]entity.Order, error) {
	rows, err := r.Db.QueryContext(ctx, Query)
	if err != nil {
		return nil, err
	}

	resp := make([]entity.Order, 0)

	for rows.Next() {
		var i entity.Order

		err := rows.Scan(
			&i.Id,
			&i.Quantity,
			&i.CustomerAddress,
		)

		if err != nil {
			return nil, err
		}

		resp = append(resp, i)
	}

	return resp, err
}
