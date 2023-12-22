package controllers

import (
	"bitbucket.org/elevatt/sirius/internal/data/contracts"
	"bitbucket.org/elevatt/sirius/internal/data/services"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	service services.AuthService
}

func GenerateAuthController(service services.AuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}

func (c *AuthController) SignInCLIENT(ctx *fiber.Ctx) error {

	body := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	err := ctx.BodyParser(&body)

	if err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	exec, err := c.service.SignInCLIENT(ctx, body.Email, body.Password)

	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(exec)

}

func (c *AuthController) SignUpCLIENT(ctx *fiber.Ctx) error {

	body := contracts.NewUserInputContract{}

	err := ctx.BodyParser(&body)

	if err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	exec, err := c.service.SignUpCLIENT(ctx, body)

	if err != nil {
		return ctx.Status(400).JSON("User already exists")
	}

	if !exec {
		return ctx.Status(400).JSON("User already exists")
	}

	return ctx.SendStatus(200)
}

func (c *AuthController) SignInADMIN(ctx *fiber.Ctx) error {
	body := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	err := ctx.BodyParser(&body)

	if err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	exec, err := c.service.SignInADMIN(ctx, body.Email, body.Password)

	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(exec)
}

func (c *AuthController) FirstAdmin(ctx *fiber.Ctx) error {

	exec, err := c.service.FirstAdmin(ctx.Context())

	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(exec)
}

func (c *AuthController) SignInWORKER(ctx *fiber.Ctx) error {

	body := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	err := ctx.BodyParser(&body)

	if err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	exec, err := c.service.SignInWORKER(ctx, body.Email, body.Password)

	if err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	return ctx.Status(200).JSON(exec)
}

func (c *AuthController) RefreshTokenADMIN(ctx *fiber.Ctx) error {

	cookie := ctx.Cookies("rt")

	exec, err := c.service.RefreshTokenADMIN(ctx, cookie)

	if err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	return ctx.Status(200).JSON(exec)
}

func (c *AuthController) RefreshTokenCLIENT(ctx *fiber.Ctx) error {

	header := ctx.Get("Authorization")

	exec, err := c.service.RefreshTokenCLIENT(ctx, header)

	if err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	return ctx.Status(200).JSON(exec)
}

func (c *AuthController) RefreshTokenWORKER(ctx *fiber.Ctx) error {

	header := ctx.Get("Authorization")

	exec, err := c.service.RefreshTokenWORKER(ctx, header)

	if err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	return ctx.Status(200).JSON(exec)
}
