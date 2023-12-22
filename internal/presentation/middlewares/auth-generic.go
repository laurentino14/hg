package middlewares

import (
	"bitbucket.org/elevatt/sirius/internal/data/utils"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func Auth(ctx *fiber.Ctx) error {

	tokenCookie := ctx.Cookies("at")
	token := strings.Replace(tokenCookie, "Bearer ", "", 1)
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	tkn, err := utils.ValidateAT(token)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	if !tkn.Valid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	claims := tkn.Claims.(*utils.Claims)

	subject, err := claims.RegisteredClaims.GetSubject()
	if err != nil {

		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	subjectJSON := &utils.Subject{}
	err = json.Unmarshal([]byte(subject), &subjectJSON)
	if err != nil {
		fmt.Println(err)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	path := ctx.Path()

	if subjectJSON.Type != 1 && strings.Contains(path, "v1/admin") {

		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	if subjectJSON.Type != 0 && strings.Contains(path, "v1/client") {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	if subjectJSON.Type != 2 && strings.Contains(path, "v1/worker") {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return ctx.Next()
}
