package domain

import "time"

type BookModel struct {
	ID          int
	Title       string
	Author      string
	Review      string
	Year        int
	Completed   bool
	AddedAt     time.Time
	CompletedAt *time.Time
}
