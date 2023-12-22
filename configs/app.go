package configs

import (
	"bitbucket.org/elevatt/sirius/configs/routes"
	"database/sql"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"log/slog"
	"os"
	"time"
)

func SiriusConfig(db *sql.DB) *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
	})

	// CORS
	//app.Use(cors.New(cors.Config{
	//	Next:             nil,
	//	AllowOriginsFunc: nil,
	//	AllowOrigins:     "*",
	//	AllowMethods:     "",
	//	AllowHeaders:     "",
	//	AllowCredentials: true,
	//	ExposeHeaders:    "",
	//	MaxAge:           0,
	//}))

	// Logger
	l := logger.New(logger.Config{
		//Format:     "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeZone:   "America/Sao_Paulo",
		TimeFormat: "02/Jan/2006 15:04:05",
		Output:     os.Stdout,
	})
	app.Use(l)

	// Compress
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	// Monitor
	if os.Getenv("ENV") == "development" {
		app.Get("/", monitor.New(monitor.Config{
			Refresh: time.Duration(1 * time.Second),
		}))
	}

	// Cache
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.QueryBool("noCache") == true
		},
		Expiration:   30 * time.Minute,
		CacheControl: true,
		CacheHeader:  "X-Cache",
		Methods:      []string{fiber.MethodGet, fiber.MethodHead},
	}))

	// ETag
	app.Use(etag.New(etag.Config{
		Weak: true,
	}))

	// Docs
	//docfilepath := path.Join(os.Getenv("PWD"), "internal/release", "docs", "swagger.json")
	//cfg := swagger.Config{
	//	BasePath: "/docs",
	//	FilePath: docfilepath,
	//	Path:     "/",
	//	Title:    "Swagger API Docs",
	//}
	//app.Use(swagger.New(cfg))

	// Routes
	v1 := app.Group("/v1").Name("v1")
	routes.SetupAdminRoutes(v1, db)

	// Hooks
	app.Hooks().OnListen(func(addr fiber.ListenData) error {
		if fiber.IsChild() {
			return nil
		}
		if os.Getenv("ENV") == "development" {
			addr.Host = "localhost"
			slog.Info("Listening on " + "http://" + addr.Host + ":" + addr.Port)
		} else {
			slog.Info("Listening on " + "http://" + addr.Host + ":" + addr.Port)
		}

		// check this type assertion to avoid a panic
		return nil
	})

	app.Hooks().OnShutdown(func() error {
		slog.Info("Shutdown App")
		return nil
	})

	return app
}
