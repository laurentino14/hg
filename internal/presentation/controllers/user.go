package controllers

import (
	"bitbucket.org/elevatt/sirius/internal/data/contracts"
	"bitbucket.org/elevatt/sirius/internal/data/services"
	"github.com/gofiber/fiber/v2"
	"log/slog"
)

type UserController struct {
	service services.UserService
}

func GenerateUserController(service services.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) Create(ctx *fiber.Ctx) error {
	body := new(contracts.NewUserInputContract)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	data, err := c.service.Create(ctx.Context(), body)
	if err != nil {
		slog.Error(err.Error())
		return ctx.Status(400).SendString(err.Error())
	}
	return ctx.Status(201).JSON(data)
}

func (c *UserController) FindAllClients(ctx *fiber.Ctx) error {
	data, err := c.service.FindAllClients(ctx.Context())
	if err != nil {
		slog.Error(err.Error())
		return ctx.Status(400).SendString(err.Error())
	}
	return ctx.Status(200).JSON(data)
}

func (c *UserController) FindAllEmployees(ctx *fiber.Ctx) error {
	data, err := c.service.FindAllEmployees(ctx.Context())
	if err != nil {
		slog.Error(err.Error())
		return ctx.Status(400).SendString(err.Error())
	}
	return ctx.Status(200).JSON(data)
}

func (c *UserController) FindAllAdmins(ctx *fiber.Ctx) error {
	data, err := c.service.FindAllAdmins(ctx.Context())
	if err != nil {
		slog.Error(err.Error())
		return ctx.Status(400).SendString(err.Error())
	}
	return ctx.Status(200).JSON(data)
}

func (c *UserController) FindByID(ctx *fiber.Ctx) error {
	data, err := c.service.FindByID(ctx.Context(), ctx.Params("id"))
	if err != nil {
		slog.Error(err.Error())
		return ctx.Status(400).SendString(err.Error())
	}

	return ctx.Status(200).JSON(data)
}

func (c *UserController) Update(ctx *fiber.Ctx) error {
	body := new(contracts.NewUserInputContract)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	data, err := c.service.Update(ctx.Context(), body)
	if err != nil {
		slog.Error(err.Error())
		return ctx.Status(400).SendString(err.Error())
	}
	return ctx.Status(200).JSON(data)
}

func (c *UserController) UpdatePassword(ctx *fiber.Ctx) error {
	body := new(contracts.UpdateUserPasswordContract)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	data, err := c.service.UpdatePassword(ctx.Context(), body)
	if err != nil {
		slog.Error(err.Error())
		return ctx.Status(400).SendString(err.Error())
	}
	return ctx.Status(200).JSON(data)
}

func (c *UserController) Delete(ctx *fiber.Ctx) error {
	return c.service.Delete(ctx.Context(), ctx.Params("id"))
}
