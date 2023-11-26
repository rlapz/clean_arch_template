package instance

import (
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rlapz/clean_arch_template/src/config"
	"github.com/rlapz/clean_arch_template/src/controller/http_controller"
	"github.com/rlapz/clean_arch_template/src/controller/http_controller/routes"
	"github.com/rlapz/clean_arch_template/src/repo"
	"github.com/rlapz/clean_arch_template/src/usecase"
	"github.com/sirupsen/logrus"
)

func setupRoutes(app *fiber.App, db *sql.DB, logger *logrus.Logger, validate *validator.Validate) {
	/*
	 * repos
	 */
	userRepo := repo.NewUserRepo(logger, db)

	/*
	 * usecases
	 */
	userUsecase := usecase.NewUserUsecase(logger, validate, userRepo)

	route := routes.RouteV1{
		App: app,

		/*
		 * controllers
		 */
		HealthController: http_controller.NewHealthController(logger),
		UserController:   http_controller.NewUserController(logger, userUsecase),
	}

	route.SetupRoutes()
}

func RunApp(isProduction bool) error {
	logger := logrus.New()
	validate := validator.New()

	config, err := config.Load(isProduction)
	if err != nil {
		return fmt.Errorf("RunApp: config.Load: %s", err)
	}

	db, err := NewDatabase(config)
	if err != nil {
		return fmt.Errorf("RunApp: NewDatabase: %s", err)
	}
	//defer db.Close()

	app := NewFiberApp(config)

	// set the routes
	setupRoutes(app, db, logger, validate)

	// gracefull shutdown
	notifChan := make(chan os.Signal, 1)
	signal.Notify(notifChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	go func() {
		s := <-notifChan

		fmt.Println()
		logger.Infof("[%+v]: Shutting down the app...", s)
		app.ShutdownWithTimeout(time.Second * 10)
	}()

	err = app.Listen(fmt.Sprintf("%s:%d", config.Http.Host, config.Http.Port))
	if err != nil {
		return fmt.Errorf("RunApp: app.Listen: %s", err)
	}

	return nil
}
