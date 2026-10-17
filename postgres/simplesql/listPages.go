package simplesql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/k0kubun/pp/v3"
)

func ListPages(
	ctx context.Context,
	conn *pgx.Conn,
	n int,
) error {
	if n <= 0 {
		return fmt.Errorf("page size must be greater than 0")
	}

	page := 1
	offset := 0

	for {
		books, err := SelectPage(ctx, conn, n, offset)
		if err != nil {
			return fmt.Errorf("get page %d: %w", page, err)
		}

		if len(books) == 0 {
			break
		}

		// fmt.Printf("Страница %d: %+v\n", page, books)
		fmt.Printf("Страница %d:\n", page)
		pp.Println(books)

		page++
		offset += n
	}

	return nil
}

func SelectPage(
	ctx context.Context,
	conn *pgx.Conn,
	limit int,
	offset int,
) ([]BookModel, error) {
	sqlQuery := `
	SELECT 
		id, 
		title, 
		author, 
		review, 
		year, 
		completed, 
		added_at, 
		completed_at
	FROM books
	ORDER BY id ASC
	LIMIT $1
	OFFSET $2;
	`

	rows, err := conn.Query(ctx, sqlQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := make([]BookModel, 0, limit)

	for rows.Next() {
		var book BookModel

		if err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Review,
			&book.Year,
			&book.Completed,
			&book.AddedAt,
			&book.CompletedAt,
		); err != nil {
			return nil, fmt.Errorf("scan book: %w", err)
		}

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate books: %w", err)
	}

	return books, nil
}
