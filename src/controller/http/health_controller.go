package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rlapz/clean_arch_template/src/model"
)

type HealthController struct {
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (u *HealthController) Check(ctx *fiber.Ctx) error {
	return ctx.JSON(
		model.WebResponse[any]{
			Success: true,
			Message: "healthy enough!",
		},
	)
}
