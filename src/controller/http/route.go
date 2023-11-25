package http

import (
	"github.com/gofiber/fiber/v2"
)

type Route struct {
	Fiber *fiber.App

	HealthController *HealthController
	UserController   *UserController
}

func (r *Route) SetupRoutes() {
	r.setupTestsV1()
	r.setupRouteGuestsV1()
}

/*
 * Version 1
 */
func (r *Route) setupRouteGuestsV1() {
	r.Fiber.Post("/api/v1/users/login", r.UserController.Login)
}

func (r *Route) setupTestsV1() {
	r.Fiber.Get("/api/v1/health", r.HealthController.Check)
}
