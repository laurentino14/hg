package controllers

import (
	"bitbucket.org/elevatt/sirius/internal/data/contracts"
	"bitbucket.org/elevatt/sirius/internal/data/services"
	"github.com/gofiber/fiber/v2"
)

type ProjectController struct {
	service services.ProjectService
}

func GenerateProjectController(service services.ProjectService) *ProjectController {
	return &ProjectController{service: service}
}

func (c *ProjectController) Create(ctx *fiber.Ctx) error {
	body := new(contracts.NewProjectInputContract)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	data, err := c.service.Create(ctx.Context(), body)
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	return ctx.Status(201).JSON(data)
}

func (c *ProjectController) FindAll(ctx *fiber.Ctx) error {
	data, err := c.service.FindAll(ctx.Context())
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	return ctx.Status(200).JSON(data)
}

func (c *ProjectController) FindAllByClientID(ctx *fiber.Ctx) error {
	data, err := c.service.FindAllByClientID(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	return ctx.Status(200).JSON(data)
}

func (c *ProjectController) FindAllByWorkerID(ctx *fiber.Ctx) error {
	data, err := c.service.FindAllByWorkerID(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	return ctx.Status(200).JSON(data)
}

func (c *ProjectController) FindByID(ctx *fiber.Ctx) error {
	data, err := c.service.FindByID(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	return ctx.Status(200).JSON(data)
}

func (c *ProjectController) Update(ctx *fiber.Ctx) error {
	body := new(contracts.UpdateProjectInputContract)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	data, err := c.service.Update(ctx.Context(), body)
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	return ctx.Status(200).JSON(data)
}

func (c *ProjectController) Delete(ctx *fiber.Ctx) error {
	err := c.service.Delete(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	return ctx.SendStatus(200)
}
