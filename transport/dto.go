package transport

import "time"

type CreateBookRequest struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Review    string `json:"review"`
	Year      int    `json:"year"`
	Completed bool   `json:"completed"`
}

type UpdateBookRequest struct {
	Title       string     `json:"title"`
	Author      string     `json:"author"`
	Review      string     `json:"review"`
	Completed   bool       `json:"completed"`
	CompletedAt *time.Time `json:"completed_at"`
}

type BookResponse struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Author      string     `json:"author"`
	Review      string     `json:"review"`
	Year        int        `json:"year"`
	Completed   bool       `json:"completed"`
	AddedAt     time.Time  `json:"added_at"`
	CompletedAt *time.Time `json:"completed_at"`
}
