package connection

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectDB(ctx context.Context) (*pgx.Conn, error) {
	connString := os.Getenv("CONN_STRING")
	return pgx.Connect(ctx, connString)
}
