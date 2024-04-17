package handlers

import (
	"github.com/Simplyphotons/fyp.git/db"
	"github.com/Simplyphotons/fyp.git/model"
	"github.com/Simplyphotons/fyp.git/security"
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
func (c Controller) GetHasProjectStatusHandler(ctx *fiber.Ctx) error {

	var (
		authority security.Authority
		ok        bool
	)
	if authority, ok = ctx.UserContext().Value(security.AuthorityKey{}).(security.Authority); !ok {
		message := model.ErrorMessage{
			Message: "cannot extract user id",
		}

		return ctx.Status(401).JSON(message)
	}

	response, err := c.dbClient.GetHasProjectStatus(ctx.Context(), authority.UserID)
	if err != nil {
		message := model.ErrorMessage{
			Message: err.Error(),
		}
		return ctx.Status(500).JSON(message)
	}
	return ctx.Status(200).JSON(response)

}

func (c Controller) NewQuestion(ctx *fiber.Ctx) error {

	err := c.dbClient.NewQuestion(ctx.Context(), db.Question{})
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
