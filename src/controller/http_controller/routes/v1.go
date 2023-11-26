package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rlapz/clean_arch_template/src/controller/http_controller"
)

/*
 * Version 1
 */
type RouteV1 struct {
	App *fiber.App

	HealthController *http_controller.HealthController
	UserController   *http_controller.UserController
}

func (r *RouteV1) SetupRoutes() {
	r.setupTests()
	r.setupRouteGuests()
}

func (r *RouteV1) setupRouteGuests() {
	r.App.Post("/api/v1/users/login", r.UserController.Login)
}

func (r *RouteV1) setupTests() {
	r.App.Get("/api/v1/health", r.HealthController.Check)
}
