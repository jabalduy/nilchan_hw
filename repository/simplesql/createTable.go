package simplesql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(
	ctx context.Context,
	conn *pgx.Conn,
) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS books (
		id SERIAL PRIMARY KEY,
		title VARCHAR(100) NOT NULL,
		author VARCHAR(100) NOT NULL,
		review VARCHAR(1000),
		year INT NOT NULL,
		completed BOOLEAN,
		added_at TIMESTAMP NOT NULL,
		completed_at TIMESTAMP
	);
	`

	_, err := conn.Exec(ctx, sqlQuery)
	return err
}
