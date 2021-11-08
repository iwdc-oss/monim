package domain

import "time"

type MockInterviewPost struct {
	ID          int32
	Name        string
	Role        string
	Logo        string
	Description string
	Agreement   string
	PublishedAt time.Time
}
