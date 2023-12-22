package controllers

import (
	"bitbucket.org/elevatt/sirius/internal/data/contracts"
	"bitbucket.org/elevatt/sirius/internal/data/services"
	"github.com/gofiber/fiber/v2"
)

type OsController struct {
	service services.OsService
}

func GenerateOsController(service services.OsService) *OsController {
	return &OsController{service: service}
}

func (c *OsController) Create(ctx *fiber.Ctx) error {
	body := new(contracts.NewOsInputContract)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	data, err := c.service.Create(ctx.Context(), body)
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	return ctx.Status(201).JSON(data)
}

func (c *OsController) FindAll(ctx *fiber.Ctx) error {
	data, err := c.service.FindAll(ctx.Context())
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	return ctx.Status(200).JSON(data)
}

func (c *OsController) FindAllByClientID(ctx *fiber.Ctx) error {
	data, err := c.service.FindAllByClientID(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	return ctx.Status(200).JSON(data)
}

func (c *OsController) FindAllByWorkerID(ctx *fiber.Ctx) error {
	data, err := c.service.FindAllByWorkerID(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	return ctx.Status(200).JSON(data)
}

func (c *OsController) FindAllByProjectCOD(ctx *fiber.Ctx) error {
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

func (c *OsController) FindByCOD(ctx *fiber.Ctx) error {
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

func (c *OsController) Update(ctx *fiber.Ctx) error {
	body := new(contracts.UpdateOsInputContract)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	data, err := c.service.Update(ctx.Context(), body)
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	return ctx.Status(200).JSON(data)
}

func (c *OsController) Delete(ctx *fiber.Ctx) error {
	cod, err := ctx.ParamsInt("cod")
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	err = c.service.Delete(ctx.Context(), cod)
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	return ctx.Status(200).JSON("Deleted")
}
