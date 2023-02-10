package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// Postgres -.
type Postgres struct {
	Pool *pgxpool.Pool
}

func New(uri string) (*Postgres, error) {

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()

	return &Postgres{}, nil
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
