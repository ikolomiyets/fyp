package db

import (
	"FYP/model"
	"context"
	"log"
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

func (db Client) GetGantt(ctx context.Context, projectIdentifier string) ([]model.Gantt, error) { //gets all milestones within a project
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

func (db Client) GetApplications(ctx context.Context) ([]model.ApplicationData, error) {
	rows, err := db.conn.QueryContext(ctx, "SELECT * from applications")
	if err != nil {
		log.Printf("cannot execute query to get applications: %v", err)
		return nil, err
	}
	result := []model.ApplicationData{}
	var (
		id           string
		studentID    string
		supervisorID string
		heading      string
		description  string
		accepted     bool
		declined     bool
	)
	for rows.Next() {
		err = rows.Scan(&id, &studentID, &supervisorID, &heading, &description, &accepted, &declined)
		if err != nil {
			log.Printf("cannot read data while getting questions: %v", err)
		}

		result = append(result, model.ApplicationData{
			ID:           id,
			StudentID:    studentID,
			SupervisorID: supervisorID,
			Heading:      heading,
			Description:  description,
			Accepted:     accepted,
			Declined:     declined,
		})
	}
	return result, nil

}

func (db Client) GetSpecificApplications(ctx context.Context, appID string) ([]model.ApplicationData, error) {
	rows, err := db.conn.QueryContext(ctx, "SELECT * from applications")
	if err != nil {
		log.Printf("cannot execute query to get applications: %v", err)
		return nil, err
	}
	result := []model.ApplicationData{}
	var (
		id           string
		studentID    string
		supervisorID string
		heading      string
		description  string
		accepted     bool
		declined     bool
	)
	for rows.Next() {
		err = rows.Scan(&id, &studentID, &supervisorID, &heading, &description, &accepted, &declined)
		if err != nil {
			log.Printf("cannot read data while getting questions: %v", err)
		}
		if id == appID {
			result = append(result, model.ApplicationData{
				ID:           id,
				StudentID:    studentID,
				SupervisorID: supervisorID,
				Heading:      heading,
				Description:  description,
				Accepted:     accepted,
				Declined:     declined,
			})
		}
	}
	return result, nil

}

func (db Client) NewQuestion(ctx context.Context) error { //adds new question to db
	//create question item for db
	//variables that are created and not taken: ID, Isanswered will be false
	return nil
}

func (db Client) NewAnswer(ctx context.Context) error { //adds new question to db
	//todo
	var questionData model.Question

	updateQuery := "UPDATE tickets SET answer = ?, is_answered = ? WHERE ticket_id = ?"

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
	id := project.ID
	studentID := project.StudentID
	name := project.Name
	time := project.CreatedAt
	supervisorID := project.SupervisorID

	updateQuery := "INSERT INTO projects (id, studentID, name, time, supervisor) VALUES (?, ?, ?, ?, ?)"

	result, err := db.conn.Exec(updateQuery, id, studentID, name, time, supervisorID)
	if err != nil {
		log.Printf("failed to add answer to corresponding ticket in db")
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	log.Printf("created %d row.\n", rowsAffected)
	return nil

}

func (db Client) CreateApplication(ctx context.Context, application Application) error {
	id := application.ID
	studentID := application.StudentID
	supervisorID := application.SupervisorID
	heading := application.Heading
	description := application.Description

	updateQuery := "INSERT INTO applications (id, studentID, supervisor, heading, description) VALUES (?, ?, ?, ?, ?)"

	result, err := db.conn.Exec(updateQuery, id, studentID, supervisorID, heading, description)
	if err != nil {
		log.Printf("failed to add answer to corresponding ticket in db")
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	log.Printf("created %d row.\n", rowsAffected)
	return nil
}

func (db Client) AcceptApplication(ctx context.Context, application Application) error {

	updateQuery := "UPDATE tickets SET accepted = ? WHERE id = ?"

	result, err := db.conn.Exec(updateQuery, true, application.ID)
	if err != nil {
		log.Printf("failed to accept application")
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	log.Printf("created %d row.\n", rowsAffected)
	return nil
}

func (db Client) DeclineApplication(ctx context.Context, application Application) error {

	updateQuery := "UPDATE tickets SET declined = ? WHERE id = ?"

	result, err := db.conn.Exec(updateQuery, true, application.ID)
	if err != nil {
		log.Printf("failed to decline application")
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	log.Printf("created %d row.\n", rowsAffected)
	return nil
}

func (db Client) CreateGanttItem(ctx context.Context, gantt Gantt) error {

	id := gantt.Id
	projectID := gantt.ProjectID
	description := gantt.Description
	startDate := gantt.StartDate
	endDate := gantt.EndDate
	feedback := gantt.Feedback
	links := gantt.Links

	updateQuery := "INSERT INTO ganttItems (id, projectID, description, startDate, endDate, feedback, links) VALUES (?, ?, ?, ?, ?, ?, ?)"

	result, err := db.conn.Exec(updateQuery, id, projectID, description, startDate, endDate, feedback, links)
	if err != nil {
		log.Printf("failed to create new gantt item/milestone to the project")
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	log.Printf("created %d row.\n", rowsAffected)
	return nil
}

func (db Client) UpdateFeedback(ctx context.Context, gantt Gantt, id string, feedback string) error {
	updateQuery := "UPDATE ganttItems SET feedback = ? WHERE id = ?"

	newFeedback := gantt.Feedback + "\n\n" + feedback

	result, err := db.conn.Exec(updateQuery, newFeedback, id)
	if err != nil {
		log.Printf("failed to update feeback")
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	log.Printf("created %d row.\n", rowsAffected)
	return nil

}
