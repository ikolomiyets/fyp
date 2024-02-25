package handlers

import (
	"FYP/model"
	"github.com/gofiber/fiber/v2"
)

func (c Controller) GetApplicationsHandler(ctx *fiber.Ctx) error { //get all gantt item for a particular project
	response, err := c.dbClient.GetApplications(ctx.Context())
	if err != nil {
		message := model.ErrorMessage{
			Message: err.Error(),
		}
		return ctx.Status(500).JSON(message)
	}
	return ctx.Status(200).JSON(response)
}

func (c Controller) GetSpecificApplicationsHandler(ctx *fiber.Ctx) error { //get all gantt item for a particular project
	response, err := c.dbClient.GetSpecificApplications(ctx.Context(), ctx.Params("id"))
	if err != nil {
		message := model.ErrorMessage{
			Message: err.Error(),
		}
		return ctx.Status(500).JSON(message)
	}
	return ctx.Status(200).JSON(response)
}
