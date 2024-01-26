package handlers

import (
	"FYP/db"
	"FYP/model"
	"context"
)

type DBClient interface {
	GetGanttItem(ctx context.Context, milestoneIdentifier string) ([]model.Gantt, error)
	CreateProject(ctx context.Context, project db.Project) error
}

type Controller struct {
	dbClient DBClient
}

func New(client DBClient) *Controller {
	return &Controller{
		dbClient: client,
	}
}
