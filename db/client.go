package db

import (
	"FYP/model"
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

func (db Client) GetQuestions(ctx context.Context) ([]model.Question, error) {
	rows, err := db.conn.QueryContext(ctx, "SELECT * from QUESTIONS")
	if err != nil {
		log.Printf("cannot execute query to get questions: %v", err)
		return nil, err
	}

	result := []model.Question{}

	var (
		id       string
		question string
	)
	for rows.Next() {
		err = rows.Scan(&id, &question)
		if err != nil {
			log.Printf("cannot read data while getting questions: %v", err)
		}

		result = append(result, model.Question{
			ID:       id,
			Question: question,
		})
	}

	// Nullable fields
	//var (
	//	id       string
	//	question sql.NullString
	//)
	//for rows.Next() {
	//	err = rows.Scan(&id, &question)
	//	if err != nil {
	//		log.Printf("cannot read data while getting questions: %v", err)
	//	}
	//
	//	item := model.Question{
	//		ID: id,
	//	}
	//
	//	if question.Valid {
	//		item.Question = &question.String
	//	}
	//	temp = append(temp)
	//}

	// Not nullable fields
	//for rows.Next() {
	//	item := model.Question{}
	//
	//	err := rows.Scan(&item.ID, &item.Question)
	//	if err != nil {
	//		log.Printf("cannot read data while getting questions: %v", err)
	//	}
	//	temp = append(temp, item)
	//}
	return result, nil

}

func (db Client) GetGantt(ctx context.Context, c *fiber.Ctx) ([]model.Gantt, error) { //gets all milestones within a project
	rows, err := db.conn.QueryContext(ctx, "SELECT * from gantt_items")
	if err != nil {
		log.Printf("cannot execute query to get questions: %v", err)
		return nil, err
	}

	result := []model.Gantt{}
	projectIdentifier := c.Params("id")
	var (
		id          string
		projectID   string
		startDate   string
		endDate     string
		description string
		links       string
		feedback    string
	)
	for rows.Next() {
		err = rows.Scan(&id, &projectID, &startDate, &endDate, &description, &links, &feedback)
		if err != nil {
			log.Printf("cannot read data while getting questions: %v", err)
		}
		if projectID == projectIdentifier {
			result = append(result, model.Gantt{
				ID:          id,
				ProjectID:   projectID,
				StartDate:   startDate,
				EndDate:     endDate,
				Description: description,
				Links:       links,
				Feedback:    feedback,
			})
		}
	}
	return result, nil
}

func (db Client) GetGanttItem(ctx context.Context, milestoneIdentifier string) ([]model.Gantt, error) { //gets one milestone
	rows, err := db.conn.QueryContext(ctx, "SELECT * from gantt_items")
	if err != nil {
		log.Printf("cannot execute query to get questions: %v", err)
		return nil, err
	}

	result := []model.Gantt{}
	var (
		id          string
		projectID   string
		startDate   string
		endDate     string
		description string
		links       string
		feedback    string
	)
	for rows.Next() {
		err = rows.Scan(&id, &projectID, &startDate, &endDate, &description, &links, &feedback)
		if err != nil {
			log.Printf("cannot read data while getting questions: %v", err)
		}
		if projectID == milestoneIdentifier {
			result = append(result, model.Gantt{
				ID:          id,
				ProjectID:   projectID,
				StartDate:   startDate,
				EndDate:     endDate,
				Description: description,
				Links:       links,
				Feedback:    feedback,
			})
		}
	}
	return result, nil
}

func (db Client) GetSupervisors(ctx context.Context) ([]model.AccountSupervisor, error) { //for use in displaying all available supervisors when a student is creating a new project application.
	rows, err := db.conn.QueryContext(ctx, "SELECT * from users")
	if err != nil {
		log.Printf("cannot execute query to get users: %v", err)
		return nil, err
	}

	result := []model.AccountSupervisor{}

	var (
		id           string
		isSupervisor bool
	)
	for rows.Next() {
		err = rows.Scan(&id, &isSupervisor)
		if err != nil {
			log.Printf("cannot read data while getting users: %v", err)
		}

		result = append(result, model.AccountSupervisor{
			ID:           id,
			IsSupervisor: isSupervisor,
		})
	}
	return result, nil
}

func (db Client) NewQuestion(ctx context.Context, c *fiber.Ctx) error { //adds new question to db
	var questionData model.Question
	if err := c.BodyParser(&questionData); err != nil {
		log.Println("Error parsing JSON:", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	id := questionData.ID
	studentID := questionData.StudentID
	supervisorID := questionData.SupervisorID
	questionShort := questionData.QuestionShort
	questionLong := questionData.Question

	temp, err := db.conn.Prepare("INSERT INTO ticket (ticket_id, student_id, supervisor_id, questionShort, questionLong, answer, isAnswered) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Printf("couldn't prepare sql query")
		return err
	}
	result, err := temp.Exec(id, studentID, supervisorID, questionShort, questionLong, "", false)
	if err != nil {
		log.Printf("couldn't fill out new entry through executing previously prepared sql query")
		return err
	}
	log.Printf("student: %v added new question %v "+studentID, result)

	return nil
}

func (db Client) NewAnswer(ctx context.Context, c *fiber.Ctx) error { //adds new question to db
	var questionData model.Question

	updateQuery := "UPDATE tickets SET answer = ?, is = is_answered = ? WHERE ticket_id = ?"

	c.Body()
	id := questionData.ID
	answer := questionData.Answer
	isAnswered := true

	result, err := db.conn.Exec(updateQuery, answer, isAnswered, id)
	if err != nil {
		log.Printf("failed to add answer to corresponding ticket in db")
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	log.Printf("Updated %d rows.\n", rowsAffected)
	return nil
}

func (db Client) CreateProject(ctx context.Context, project Project) error {
	return nil
}
