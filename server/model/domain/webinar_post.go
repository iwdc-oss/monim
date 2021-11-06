package domain

import "time"

type WebinarPost struct {
	Id          int32
	Name        string
	Banner      string
	Description string
	PublishedAt time.Time
}
