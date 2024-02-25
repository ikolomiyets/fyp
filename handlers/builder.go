package handlers

import (
	"FYP/db"
	"FYP/model"
	"context"
)

type DBClient interface {
	GetGanttItem(ctx context.Context, milestoneIdentifier string) ([]model.Gantt, error)
	CreateProject(ctx context.Context, project db.Project) error
	GetQuestions(ctx context.Context) ([]model.Question, error)
	GetSupervisors(ctx context.Context) ([]model.AccountSupervisor, error)
	GetApplications(ctx context.Context) ([]model.ApplicationData, error)
	GetSpecificApplications(ctx context.Context, appID string) ([]model.ApplicationData, error)
	NewQuestion(ctx context.Context) error
	NewAnswer(ctx context.Context) error
	GetGantt(ctx context.Context, projectIdentifier string) ([]model.Gantt, error)
}

type Controller struct {
	dbClient DBClient
}

func New(client DBClient) *Controller {
	return &Controller{
		dbClient: client,
	}
}
