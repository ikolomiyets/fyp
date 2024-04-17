package handlers

import (
	"context"
	"github.com/Simplyphotons/fyp.git/db"
	"github.com/Simplyphotons/fyp.git/model"
	"github.com/gofiber/fiber/v2"
	"testing"
)

type DBMock struct {
	GetGanttItemResponse   []model.Gantt
	GetGanttItemError      error
	GetGanttItemCallNumber int
}

func (db DBMock) GetProjects(ctx context.Context, supervisor_id string) ([]model.ProjectData, error) {
	//TODO implement me
	panic("implement me")
}

func (db DBMock) GetUsername(ctx context.Context, userId string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (db DBMock) CreateGanttItem(ctx context.Context, gantt db.Gantt) error {
	//TODO implement me
	panic("implement me")
}

func (db DBMock) UpdateFeedback(ctx context.Context, gantt db.Gantt, id string, feedback string) error {
	//TODO implement me
	panic("implement me")
}

func (db DBMock) CreateApplication(ctx context.Context, application db.Application) error {
	//TODO implement me
	panic("implement me")
}

func (db DBMock) AcceptApplication(ctx context.Context, application db.Application) error {
	//TODO implement me
	panic("implement me")
}

func (db DBMock) DeclineApplication(ctx context.Context, application db.Application) error {
	//TODO implement me
	panic("implement me")
}

func (db DBMock) GetQuestions(ctx context.Context) ([]model.Question, error) {
	//TODO implement me
	panic("implement me")
}

func (db DBMock) GetSupervisors(ctx context.Context) ([]model.UserData, error) {
	//TODO implement me
	panic("implement me")
}

func (db DBMock) GetApplications(ctx context.Context) ([]model.ApplicationData, error) {
	//TODO implement me
	panic("implement me")
}

func (db DBMock) GetSpecificApplications(ctx context.Context, appID string) ([]model.ApplicationData, error) {
	//TODO implement me
	panic("implement me")
}

func (db DBMock) NewQuestion(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (db DBMock) NewAnswer(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (db DBMock) GetGantt(ctx context.Context, projectIdentifier string) ([]model.Gantt, error) {
	//TODO implement me
	panic("implement me")
}

func (db DBMock) GetGanttItem(ctx context.Context, milestoneIdentifier string) ([]model.Gantt, error) {
	db.GetGanttItemCallNumber++
	return db.GetGanttItemResponse, db.GetGanttItemError
}

func (db DBMock) CreateProject(ctx context.Context, project db.Project) error {
	return nil
}

func TestGetGanttItem(t *testing.T) {
	dbMock := DBMock{
		GetGanttItemResponse: []model.Gantt{
			{
				ID:        "bc11d336-241d-4d69-8061-bfca6e39809e",
				ProjectID: "1234",

				StartDate:   "01/24",
				EndDate:     "02/24",
				Description: "description text",
			},
			{
				ID: "76906af6-26b2-4f9e-8c78-7201807f6a2b",
			},
		},
	}

	testController := New(dbMock)

	ctx := &fiber.Ctx{}

	err := testController.GetGantt(ctx)
	if err != nil {
		t.Error(err)
	}
}
