package handlers

import (
	"FYP/db"
	"FYP/model"
	"FYP/security"
	"encoding/json"
	"fmt"
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

func (c Controller) CreateApplicationHandler(ctx *fiber.Ctx) error {
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

	fmt.Printf("%s", authority.UserID)

	// Read the request body
	var application model.ApplicationData

	err := json.Unmarshal(ctx.Body(), &application)
	if err != nil {
		message := model.ErrorMessage{
			Message: err.Error(),
		}
		return ctx.Status(400).JSON(message)
	}

	// Translate it to the db request
	applicationRequest := db.Application{
		ID:           application.ID,
		StudentID:    application.StudentID,
		SupervisorID: application.SupervisorID,
		Heading:      application.Heading,
		Description:  application.Description,
		Accepted:     false,
		Declined:     false,
	}

	// Execute db request
	err = c.dbClient.CreateApplication(ctx.Context(), applicationRequest)
	if err != nil {
		message := model.ErrorMessage{
			Message: err.Error(),
		}
		return ctx.Status(500).JSON(message)
	}

	return ctx.SendStatus(204)
}

func (c Controller) AcceptApplicationHandler(ctx *fiber.Ctx) error {
	var application model.ApplicationData

	err := json.Unmarshal(ctx.Body(), &application)
	if err != nil {
		message := model.ErrorMessage{
			Message: err.Error(),
		}
		return ctx.Status(400).JSON(message)
	}

	// Translate it to the db request
	applicationRequest := db.Application{
		ID:           application.ID,
		StudentID:    application.StudentID,
		SupervisorID: application.SupervisorID,
		Heading:      application.Heading,
		Description:  application.Description,
		Accepted:     false,
		Declined:     false,
	}

	// Execute db request
	err = c.dbClient.AcceptApplication(ctx.Context(), applicationRequest)
	if err != nil {
		message := model.ErrorMessage{
			Message: err.Error(),
		}
		return ctx.Status(500).JSON(message)
	}

	return ctx.SendStatus(204)
}

func (c Controller) DeclineApplicationHandler(ctx *fiber.Ctx) error {
	var application model.ApplicationData

	err := json.Unmarshal(ctx.Body(), &application)
	if err != nil {
		message := model.ErrorMessage{
			Message: err.Error(),
		}
		return ctx.Status(400).JSON(message)
	}

	// Translate it to the db request
	applicationRequest := db.Application{
		ID:           application.ID,
		StudentID:    application.StudentID,
		SupervisorID: application.SupervisorID,
		Heading:      application.Heading,
		Description:  application.Description,
		Accepted:     false,
		Declined:     false,
	}

	// Execute db request
	err = c.dbClient.DeclineApplication(ctx.Context(), applicationRequest)
	if err != nil {
		message := model.ErrorMessage{
			Message: err.Error(),
		}
		return ctx.Status(500).JSON(message)
	}

	return ctx.SendStatus(204)
}
