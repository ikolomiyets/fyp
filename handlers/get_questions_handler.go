package handlers

import (
	"FYP/model"
	"github.com/gofiber/fiber/v2"
)

func (c Controller) GetQuestionsHandler(ctx *fiber.Ctx) error {

	response := model.GetQuestionsResponse{
		Questions: []model.Question{
			{
				ID:       "bc11d336-241d-4d69-8061-bfca6e39809e",
				Question: "What is your name?",
			},
			{
				ID:       "76906af6-26b2-4f9e-8c78-7201807f6a2b",
				Question: "What is your date of birth?",
			},
		},
	}

	return ctx.Status(200).JSON(response)

}

func (c Controller) GetSupervisorHandler(ctx *fiber.Ctx) error {

	response := model.GetAccountSuperResponse{
		IsSupervisor: []model.AccountSupervisor{
			{
				ID:           "bc11d336-241d-4d69-8061-bfca6e39809e",
				IsSupervisor: true,
			},
			{
				ID:           "76906af6-26b2-4f9e-8c78-7201807f6a2b",
				IsSupervisor: false,
			},
		},
	}

	return ctx.Status(200).JSON(response)
}

func (c Controller) NewQuestion(ctx *fiber.Ctx) error {

	response := model.GetQuestionsResponse{
		Questions: []model.Question{
			{
				ID:            "bc11d336-241d-4d69-8061-bfca6e39809e",
				StudentID:     "1Student",
				SupervisorID:  "1Super-Visor",
				QuestionShort: "Why is milestone 1 x problem",
				Question:      "Big long question text,Big long question text,Big long question text,Big long question text,Big long question text,Big long question text,Big long question text,Big long question text,",
			},
		},
	}

	return ctx.Status(200).JSON(response)

}

func (c Controller) NewAnswer(ctx *fiber.Ctx) error {

	response := model.GetQuestionsResponse{
		Questions: []model.Question{
			{
				ID:     "bc11d336-241d-4d69-8061-bfca6e39809e",
				Answer: "My name is Jeff",
			},
		},
	}

	return ctx.Status(200).JSON(response)

}

func (c Controller) getGanttItem(ctx *fiber.Ctx) error {

	response := model.GetGanttResponse{
		Gantt: []model.Gantt{
			{
				ID:          "bc11d336-241d-4d69-8061-bfca6e39809e",
				ProjectID:   "1234",
				StartDate:   "01/24",
				EndDate:     "02/24",
				Description: "description text",
				Links:       "https://link.com",
				Feedback:    "I dont like it",
			},
		},
	}
	return ctx.Status(200).JSON(response)
}
