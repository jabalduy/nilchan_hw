package simplesql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func SelectRows(ctx context.Context, conn *pgx.Conn) ([]BookModel, error) {
	sqlQuery := `
	SELECT id, title, author, review, year, completed, added_at, completed_at
	FROM books
	`

	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := make([]BookModel, 0)

	for rows.Next() {
		var book BookModel

		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Review,
			&book.Year,
			&book.Completed,
			&book.AddedAt,
			&book.CompletedAt,
		)
		if err != nil {
			return nil, err
		}

		books = append(books, book)

		printBook(book)
	}

	return books, nil
}

func printBook(book BookModel) {
	fmt.Println("------------------")
	fmt.Println("id:", book.ID)
	fmt.Println("title:", book.Title)
	fmt.Println("author:", book.Author)
	fmt.Println("year:", book.Year)
	fmt.Println("completed:", book.Completed)
	fmt.Println("added at:", book.AddedAt)
	fmt.Println("completed at:", book.CompletedAt)
}
