package db

import "time"

type Project struct {
	ID           string
	Name         string
	StudentID    string
	SupervisorID string
	CreatedAt    time.Time
}
