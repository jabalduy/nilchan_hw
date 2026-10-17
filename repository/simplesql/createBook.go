package simplesql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateBook(
	ctx context.Context,
	conn *pgx.Conn,
	book BookModel,
) (int, error) {
	sqlQuery := `
	INSERT INTO books (
		title, 
		author, 
		review, 
		year, 
		completed, 
		added_at, 
		completed_at
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id
	`

	var id int

	err := conn.QueryRow(
		ctx,
		sqlQuery,
		book.Title,
		book.Author,
		book.Review,
		book.Year,
		book.Completed,
		book.AddedAt,
		book.CompletedAt,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, err
}
