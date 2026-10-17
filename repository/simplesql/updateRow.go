package simplesql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func UpdateRow(ctx context.Context, conn *pgx.Conn, book BookModel) error {
	sqlQuery := `
	UPDATE books
	SET title=$1, author=$2, review=$3, year=$4, completed=$5, added_at=$6, completed_at=$7
	WHERE id=$8
	`
	tag, err := conn.Exec(
		ctx,
		sqlQuery,
		book.Title,
		book.Author,
		book.Review,
		book.Year,
		book.Completed,
		book.AddedAt,
		book.CompletedAt,
		book.ID,
	)

	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("book with id %d not found", book.ID)
	}

	return nil
}
