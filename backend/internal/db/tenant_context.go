package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TenantConn struct {
	Pool *pgxpool.Pool
}

func (t TenantConn) WithTenant(ctx context.Context, slug string) (context.Context, error) {
	_, err := t.Pool.Exec(ctx, fmt.Sprintf(`set local search_path = "%s", public`, slug))
	return ctx, err
}

func (t TenantConn) WithPublic(ctx context.Context) (context.Context, error) {
	_, err := t.Pool.Exec(ctx, `set local search_path = public`)
	return ctx, err
}
