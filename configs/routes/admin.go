package routes

import (
	"bitbucket.org/elevatt/sirius/configs/factories"
	"bitbucket.org/elevatt/sirius/internal/presentation/middlewares"
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

func SetupAdminRoutes(router fiber.Router, db *sql.DB) {
	admin := router.Group("/admin").Name("Admin")

	// Auth
	authController := factories.GenerateAuthController(db)
	auth := admin.Group("/auth")
	auth.Post("/signin", authController.SignInADMIN).Name("Sign In ADMIN")
	auth.Get("/refresh", authController.RefreshTokenADMIN).Name("Refresh Token ADMIN")
	auth.Put("/first", authController.FirstAdmin).Name("First Admin")

	// User
	userController := factories.GenerateUserController(db)
	user := admin.Group("/user", middlewares.Auth)
	user.Post("/", userController.Create).Name("Create User")
	user.Get("/clients", userController.FindAllClients).Name("Find All Clients")
	user.Get("/employees", userController.FindAllEmployees).Name("Find All Employees")
	user.Get("/admins", userController.FindAllAdmins).Name("Find All Admins")
	user.Get("/:id", userController.FindByID).Name("Find User By ID")
	user.Patch("/:id", userController.Update).Name("Update User")
	user.Patch("/:id/password", userController.UpdatePassword).Name("Update User Password")
	user.Delete("/:id", userController.Delete).Name("Delete User")

	// Project
	projectController := factories.GenerateProjectController(db)
	project := admin.Group("/project", middlewares.Auth)
	project.Post("/", projectController.Create).Name("Create Project")
	project.Get("/", projectController.FindAll).Name("Find All Projects")
	project.Get("/client/:id", projectController.FindAllByClientID).Name("Find All Projects By Client ID")
	project.Get("/worker/:id", projectController.FindAllByWorkerID).Name("Find All Projects By Worker ID")
	project.Get("/:cod", projectController.FindByID).Name("Find Project By ID")
	project.Patch("/:cod", projectController.Update).Name("Update Project")
	project.Delete("/:cod", projectController.Delete).Name("Delete Project")

	// Os
	osController := factories.GenerateOsController(db)
	os := admin.Group("/os", middlewares.Auth)
	os.Post("/", osController.Create).Name("Create Os")
	os.Get("/", osController.FindAll).Name("Find All Os")
	os.Get("/client/:id", osController.FindAllByClientID).Name("Find All Os By Client ID")
	os.Get("/worker/:id", osController.FindAllByWorkerID).Name("Find All Os By Worker ID")
	os.Get("/project/:cod", osController.FindAllByProjectCOD).Name("Find All Os By Project COD")
	os.Get("/:cod", osController.FindByCOD).Name("Find Os By ID")
	os.Patch("/:cod", osController.Update).Name("Update Os")
	os.Delete("/:cod", osController.Delete).Name("Delete Os")

	// Task
	taskController := factories.GenerateTaskController(db)
	task := admin.Group("/task", middlewares.Auth)
	task.Post("/", taskController.Create).Name("Create Task")
	task.Get("/", taskController.FindAll).Name("Find All Task")
	task.Get("/worker/:id", taskController.FindAllByWorkerID).Name("Find All Task By Worker ID")
	task.Get("/project/:cod", taskController.FindAllByProjectCOD).Name("Find All Task By Project COD")
	task.Get("/:cod", taskController.FindByCOD).Name("Find Task By ID")
	task.Patch("/:cod", taskController.Update).Name("Update Task")
	task.Delete("/:cod", taskController.Delete).Name("Delete Task")
}
