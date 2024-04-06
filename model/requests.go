package model

import "time"

type AuthorizationRequest struct { //400
	Code         string `json:"code"`
	RefreshToken string `json:"refresh_token"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

type GetQuestionsResponse struct {
	Questions []Question `json:"questions"`
}

type Question struct {
	ID            string `json:"id"`
	StudentID     string `json:"studentID,omitempty"`
	SupervisorID  string `json:"supervisorID,omitempty"`
	QuestionShort string `json:"questionShort,omitempty"`
	Question      string `json:"question,omitempty"`
	Answer        string `json:"answer"`
	IsAnswered    bool   `json:"isAnswered"`

	//Question string `json:"question,omitempty"`
}

type GetGanttResponse struct {
	Gantt []Gantt `json:"gantt"`
}

type Gantt struct {
	ID          string `json:"id"`
	ProjectID   string `json:"projectID"`
	GanttName   string `json:"ganttName"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	Description string `json:"description"`
	Links       string `json:"links"`
	Feedback    string `json:"Feedback"`
}

type AccountSupervisor struct { //for checking account type
	ID           string `json:"id"`
	IsSupervisor bool   `json:"isSupervisor"`
}

type GetAccountSuperResponse struct {
	IsSupervisor []AccountSupervisor `json:"isSupervisors"`
}

type QuestionData struct {
	ID            string `json:"id"`
	StudentID     string `json:"studentID"`
	SupervisorID  string `json:"supervisorID"`
	QuestionShort string `json:"QuestionShort"`
	QuestionLong  string `json:"QuestionLong"`
}

type ApplicationData struct {
	ID           string `json:"id"`
	StudentID    string `json:"studentID"`
	SupervisorID string `json:"supervisorID"`
	Heading      string `json:"heading"`
	Description  string `json:"description"`
	Accepted     bool   `json:"accepted"`
	Declined     bool   `json:"declined"`
}

type ProjectData struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	StudentID    string    `json:"studentID"`
	SupervisorID string    `json:"supervisorID"`
	CreatedAt    time.Time `json:"createdAt"`
}
