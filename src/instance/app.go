package instance

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rlapz/clean_arch_template/src/config"
	"github.com/rlapz/clean_arch_template/src/controller/http"
	"github.com/rlapz/clean_arch_template/src/model"
	"github.com/rlapz/clean_arch_template/src/repo"
	"github.com/rlapz/clean_arch_template/src/usecase"
)

func NewRoutes(db *sql.DB, fiberApp *fiber.App, config *config.Config) *http.Route {
	/*
	 * repos
	 */
	userRepo := repo.NewUserRepo(config.Log, db)

	/*
	 * usecases
	 */
	userUsecase := usecase.NewUserUsecase(config.Log, config.Validate, userRepo)

	/*
	 * controllers
	 */
	healthController := http.NewHealthController(config.Log)
	userController := http.NewUserController(config.Log, userUsecase)

	return &http.Route{
		Fiber: fiberApp,

		HealthController: healthController,
		UserController:   userController,
	}
}

func RunApp(isProduction bool) error {
	config, err := config.Load(isProduction)
	if err != nil {
		return fmt.Errorf("RunApp: config.Load: %s", err)
	}

	db, err := NewDatabase(config)
	if err != nil {
		return fmt.Errorf("RunApp: NewDatabase: %s", err)
	}

	fiberApp := fiber.New(fiber.Config{
		AppName: config.AppName,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			status := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				status = e.Code
			}

			return ctx.Status(status).JSON(
				model.WebResponse[any]{
					Success: false,
					Message: err.Error(),
				},
			)
		},
		Prefork: config.Http.IsPrefork,
	})

	// set the routes
	NewRoutes(db, fiberApp, config).SetupRoutes()

	err = fiberApp.Listen(fmt.Sprintf("%s:%d", config.Http.Host, config.Http.Port))
	if err != nil {
		return fmt.Errorf("RunApp: fiberApp.Listen: %s", err)
	}

	return nil
}
