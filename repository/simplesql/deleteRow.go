package simplesql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func DeleteRow(ctx context.Context, conn *pgx.Conn, ids []int) error {
	sqlQuery := `
	DELETE FROM books
	WHERE id = ANY($1);
	`

	tag, err := conn.Exec(ctx, sqlQuery, ids)
	if err != nil {
		return fmt.Errorf("delete books: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("no books found for ids %v", ids)
	}

	return nil
}
