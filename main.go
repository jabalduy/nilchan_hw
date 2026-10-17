package main

import (
	"context"
	"fmt"
	"nilchan_hw/postgres/connection"
	"nilchan_hw/postgres/simplesql"
	"os"
	"strconv"

	"go.uber.org/zap"
)

func main() {
	strCount := os.Getenv("OUTPUT_COUNT")
	count, err := strconv.Atoi(strCount)
	if err != nil {
		panic(err)
	}

	for i := range count {
		fmt.Println("Работа цикла:", i+1)
	}

	// os.Exit(runMain())
}

func runMain() int {
	logger, logFileClose, err := NewLogger("DEBUG")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize logger: %v\n", err)
		return 1
	}
	defer logFileClose()
	defer logger.Sync()

	if err := run(logger); err != nil {
		logger.Error("application stopped with error",
			zap.Error(err),
		)
		return 1
	}

	return 0
}

func run(logger *zap.Logger) error {
	ctx := context.Background()

	logger.Info("connecting to DB")

	conn, err := connection.ConnectDB(ctx)
	if err != nil {
		return fmt.Errorf("connect to database: %w", err)
	}
	defer conn.Close(ctx)

	if err = conn.Ping(ctx); err != nil {
		return fmt.Errorf("ping database: %w", err)
	}

	logger.Info("connected to database")

	if err = simplesql.CreateTable(ctx, conn); err != nil {
		return fmt.Errorf("create books table: %w", err)
	}

	logger.Info("books table is ready")

	// ЗАПИСАТЬ
	// book := simplesql.BookModel{
	// 	Title:     "Вишневый сад",
	// 	Author:    "Чехов А. П.",
	// 	Review:    "Театральная классика",
	// 	Year:      1903,
	// 	Completed: false,
	// 	AddedAt:   time.Now(),
	// }

	// id, err := simplesql.CreateBook(ctx, conn, book)
	// if err != nil {
	// 	return fmt.Errorf("create book %q: %w", book.Title, err)
	// }

	// logger.Info("book created",
	// 	zap.Int("id", id),
	// 	zap.String("title", book.Title),
	// )

	// ОБНОВИТЬ
	// books, err := simplesql.SelectRows(ctx, conn)
	// if err != nil {
	// 	return fmt.Errorf("select rows: %w", err)
	// }

	// for _, book := range books {
	// 	if book.ID == 3 {
	// 		book.Title = "Три сестры"
	// 		book.Author = "Чехов А. П."
	// 		book.Year = 1901
	// 		book.Review = "Топчик"
	// 		book.Completed = true
	// 		now := time.Now()
	// 		book.CompletedAt = &now

	// 		if err := simplesql.UpdateRow(ctx, conn, book); err != nil {
	// 			return fmt.Errorf("update row %d: %w", book.ID, err)
	// 		}

	// 		logger.Info("book updated",
	// 			zap.Int("id", book.ID),
	// 			zap.String("title", book.Title),
	// 		)

	// 		break
	// 	}
	// }

	// УДАЛИТЬ
	// ids := []int{1, 5, 7}

	// if err = simplesql.DeleteRow(ctx, conn, ids); err != nil {
	// 	return fmt.Errorf("delete rows %v: %w", ids, err)
	// }

	// logger.Info("books deleted",
	// 	zap.Ints("ids", ids),
	// )

	// ПАГИНАЦИЯ
	if err := simplesql.ListPages(ctx, conn, 5); err != nil {
		return fmt.Errorf("list pages: %w", err)
	}

	logger.Info("books pagination completed")

	return nil
}
