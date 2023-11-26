package http_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rlapz/clean_arch_template/src/model"
	"github.com/rlapz/clean_arch_template/src/util"
)

type HealthController struct {
	log *util.Logger
}

func NewHealthController(logger *util.Logger) *HealthController {
	return &HealthController{
		log: logger,
	}
}

func (u *HealthController) Check(ctx *fiber.Ctx) error {
	u.log.Infof("hello world!")
	return ctx.JSON(
		model.HttpResponse[any]{
			Success: true,
			Message: "healthy enough!",
		},
	)
}
