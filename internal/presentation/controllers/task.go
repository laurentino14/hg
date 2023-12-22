package controllers

import (
	"bitbucket.org/elevatt/sirius/internal/data/contracts"
	"bitbucket.org/elevatt/sirius/internal/data/services"
	"github.com/gofiber/fiber/v2"
)

type TaskController struct {
	service services.TaskService
}

func GenerateTaskController(service services.TaskService) *TaskController {
	return &TaskController{service: service}
}

func (c *TaskController) Create(ctx *fiber.Ctx) error {
	body := new(contracts.NewTaskInputContract)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	data, err := c.service.Create(ctx.Context(), body)
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	return ctx.Status(201).JSON(data)
}

func (c *TaskController) FindAll(ctx *fiber.Ctx) error {
	data, err := c.service.FindAll(ctx.Context())
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	return ctx.Status(200).JSON(data)
}

func (c *TaskController) FindAllByProjectCOD(ctx *fiber.Ctx) error {
	cod, err := ctx.ParamsInt("cod")
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	data, err := c.service.FindAllByProjectCOD(ctx.Context(), cod)
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	return ctx.Status(200).JSON(data)
}

func (c *TaskController) FindAllByWorkerID(ctx *fiber.Ctx) error {
	data, err := c.service.FindAllByWorkerID(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	return ctx.Status(200).JSON(data)
}

func (c *TaskController) FindByCOD(ctx *fiber.Ctx) error {
	cod, err := ctx.ParamsInt("cod")
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	data, err := c.service.FindByCOD(ctx.Context(), cod)
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	return ctx.Status(200).JSON(data)
}

func (c *TaskController) Update(ctx *fiber.Ctx) error {
	body := new(contracts.UpdateTaskInputContract)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	data, err := c.service.Update(ctx.Context(), body)
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	return ctx.Status(200).JSON(data)
}

func (c *TaskController) Delete(ctx *fiber.Ctx) error {
	cod, err := ctx.ParamsInt("cod")
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	err = c.service.Delete(ctx.Context(), cod)
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	return ctx.Status(200).SendString("Deleted")
}
