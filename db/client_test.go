package db

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Simplyphotons/fyp.git/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_GetQuestions(t *testing.T) {
	documents := []model.Question{
		{
			ID:            "92d1cb6d-53c6-4bd0-bb10-1d64e98bfc92",
			StudentID:     "5f864cb5-a86b-4b74-aa07-44492f9cc201",
			SupervisorID:  "4c6ee9e2-c667-4190-91dc-1d6f3be85046",
			QuestionShort: "Test Question 1",
			Question:      "Question Long 1",
			Answer:        "Test Answer 1",
			IsAnswered:    false,
		},
		{
			ID:           "bfd66ee9-a04a-481d-88ef-72f2b20d2685",
			StudentID:    "002c8664-f664-45ab-ab6f-b7a8239ce49d",
			SupervisorID: "c6080ce0-0abc-4bb3-a852-92e1864e50b4",
			Question:     "Question Long 2",
			IsAnswered:   false,
		},
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT ticket_id, question from tickets").
		WillReturnRows(sqlmock.NewRows(
			[]string{
				"ticket_id", "question",
			},
		).
			AddRow(documents[0].ID, documents[0].Question).
			AddRow(documents[1].ID, documents[1].Question)).
		RowsWillBeClosed()

	d := &Client{
		conn: db,
	}

	res, err := d.GetQuestions(context.Background())
	if !assert.Nil(t, err) {
		return
	}

	assert.Equal(t, 2, len(res))
	for i, row := range res {
		assert.Equal(t, documents[i].ID, row.ID)
		assert.Equal(t, documents[i].Question, row.Question)
	}
}

func TestClient_GetQuestionsWithQueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT ticket_id, question from tickets").WillReturnError(errors.New("cannot query"))

	d := &Client{
		conn: db,
	}

	res, err := d.GetQuestions(context.Background())
	if !assert.NotNil(t, err) {
		return
	}

	assert.Equal(t, 0, len(res))
	assert.Equal(t, "cannot query", err.Error())
}

func TestClient_GetQuestionsWithScanError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT ticket_id, question from tickets").
		WillReturnRows(sqlmock.NewRows(
			[]string{
				"ticket_id", "question",
			},
		).
			AddRow(nil, "")).
		RowsWillBeClosed()

	d := &Client{
		conn: db,
	}

	res, err := d.GetQuestions(context.Background())
	if !assert.NotNil(t, err) {
		return
	}

	assert.Equal(t, 0, len(res))
	assert.Equal(t, "sql: Scan error on column index 0, name \"ticket_id\": converting NULL to string is unsupported", err.Error())
}
