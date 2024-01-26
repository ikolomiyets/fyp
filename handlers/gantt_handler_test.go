package handlers

import (
	"FYP/db"
	"FYP/model"
	"context"
	"github.com/gofiber/fiber/v2"
	"testing"
)

type DBMock struct {
	GetGanttItemResponse   []model.Gantt
	GetGanttItemError      error
	GetGanttItemCallNumber int
}

func (db DBMock) GetGanttItem(ctx context.Context, milestoneIdentifier string) ([]model.Gantt, error) {
	db.GetGanttItemCallNumber++
	return db.GetGanttItemResponse, db.GetGanttItemError
}

func (db DBMock) CreateProject(ctx context.Context, project db.Project) error {
	return nil
}

func TestGetGanttoItem(t *testing.T) {
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

	err := testController.GetGanttoItem(ctx)
	if err != nil {
		t.Error(err)
	}
}
