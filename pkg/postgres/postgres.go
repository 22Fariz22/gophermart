package postgres

import (
	"github.com/22Fariz22/gophermart/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

const (
	_defaultMaxPoolSize  = 1
	_defaultConnAttempts = 10
	_defaultConnTimeout  = time.Second
)

// Postgres -.
type Postgres struct {
	Pool *pgxpool.Pool
}

func New(cfg *config.Config) (*Postgres, error) {
	pg := &Postgres{}

	return pg, nil
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
