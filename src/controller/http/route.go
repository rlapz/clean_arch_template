package http

import (
	"github.com/gofiber/fiber/v2"
)

type Route struct {
	Fiber *fiber.App

	HealthController *HealthController
	UserController   *UserController
}

func (r *Route) SetupRoute() {
	r.setupRouteGuestV1()
}

/*
 * Version 1
 */
func (r *Route) setupRouteGuestV1() {
	r.Fiber.Get("/api/v1/health", r.HealthController.Check)
	r.Fiber.Post("/api/v1/users/login", r.UserController.Login)
}
