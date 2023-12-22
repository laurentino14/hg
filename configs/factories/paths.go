package factories

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func Paths(app *fiber.App, startPath string) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		routes := app.GetRoutes(true)
		var data []fiber.Route
		for _, route := range routes {
			if strings.Contains(route.Path, startPath) == true {
				data = append(data, route)
			}
		}
		result, _ := json.MarshalIndent(data, "", "  ")
		return ctx.Status(200).SendString(string(result))
	}
}
