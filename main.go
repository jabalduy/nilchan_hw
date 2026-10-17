package main

import (
	"context"
	"fmt"
	"nilchan_hw/postgres/connection"
)

func main() {
	fmt.Println("Hello, World!")

	ctx := context.Background()
	conn, err := connection.ConnectDB(ctx)
	if err != nil {
		panic(err)
	}

	if err = conn.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Connected seccesfully!")
}
