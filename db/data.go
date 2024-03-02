package db

import "time"

type Project struct {
	ID           string
	Name         string
	StudentID    string
	SupervisorID string
	CreatedAt    time.Time
}

type Application struct {
	ID           string
	StudentID    string
	SupervisorID string
	Heading      string
	Description  string
	Accepted     bool
	Declined     bool
}

type Gantt struct {
	Id          string
	ProjectID   string
	StartDate   string
	EndDate     string
	Description string
	Links       string
	Feedback    string
}

type Question struct {
	Id string
	//todo

}
