package handlers

import (
	"FYP/model"
	"github.com/gofiber/fiber/v2"
)

func (c Controller) GetQuestionsHandler(ctx *fiber.Ctx) error {

	response, err := c.dbClient.GetQuestions(ctx.Context())
	if err != nil {
		message := model.ErrorMessage{
			Message: err.Error(),
		}
		return ctx.Status(500).JSON(message)
	}
	return ctx.Status(200).JSON(response)

}

func (c Controller) GetSupervisorHandler(ctx *fiber.Ctx) error {

	response, err := c.dbClient.GetSupervisors(ctx.Context())
	if err != nil {
		message := model.ErrorMessage{
			Message: err.Error(),
		}
		return ctx.Status(500).JSON(message)
	}
	return ctx.Status(200).JSON(response)

}

func (c Controller) NewQuestion(ctx *fiber.Ctx) error {

	err := c.dbClient.NewQuestion(ctx.Context())
	if err != nil {
		message := model.ErrorMessage{
			Message: err.Error(),
		}
		return ctx.Status(500).JSON(message)
	}
	return ctx.Status(200).JSON("success")
}

func (c Controller) NewAnswer(ctx *fiber.Ctx) error {

	err := c.dbClient.NewAnswer(ctx.Context())
	if err != nil {
		message := model.ErrorMessage{
			Message: err.Error(),
		}
		return ctx.Status(500).JSON(message)
	}
	return ctx.Status(200).JSON("success")

}
