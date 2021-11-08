package domain

import "time"

type WebinarPost struct {
	ID          int32
	Name        string
	Banner      string
	Description string
	PublishedAt time.Time
}
