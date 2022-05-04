package models

import "time"

type Link struct {
	ID        int
	Code      string // Код сокращенного URL
	LongURL   string // Исходный URL
	ShortURL  string // Сокращённый URL
	Tags      []string
	CreatedAt time.Time
}
