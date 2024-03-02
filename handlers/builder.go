package handlers

import (
	"FYP/db"
	"FYP/model"
	"context"
)

type DBClient interface {
	GetGanttItem(ctx context.Context, milestoneIdentifier string) ([]model.Gantt, error)
	CreateProject(ctx context.Context, project db.Project) error
	CreateApplication(ctx context.Context, application db.Application) error
	AcceptApplication(ctx context.Context, application db.Application) error
	DeclineApplication(ctx context.Context, application db.Application) error
	GetQuestions(ctx context.Context) ([]model.Question, error)
	GetSupervisors(ctx context.Context) ([]model.AccountSupervisor, error)
	GetApplications(ctx context.Context) ([]model.ApplicationData, error)
	GetSpecificApplications(ctx context.Context, appID string) ([]model.ApplicationData, error)
	NewQuestion(ctx context.Context) error
	NewAnswer(ctx context.Context) error
	GetGantt(ctx context.Context, projectIdentifier string) ([]model.Gantt, error)
	CreateGanttItem(ctx context.Context, gantt db.Gantt) error
	UpdateFeedback(ctx context.Context, gantt db.Gantt, id string, feedback string) error
}

type Controller struct {
	dbClient DBClient
}

func New(client DBClient) *Controller {
	return &Controller{
		dbClient: client,
	}
}
