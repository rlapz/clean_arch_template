package http

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rlapz/clean_arch_template/src/model"
	"github.com/sirupsen/logrus"
)

type HealthController struct {
	log *logrus.Logger
}

func NewHealthController(logger *logrus.Logger) *HealthController {
	return &HealthController{
		log: logger,
	}
}

func (u *HealthController) Check(ctx *fiber.Ctx) error {
	u.log.Printf("hello world! %+v", time.Now())
	return ctx.JSON(
		model.WebResponse[any]{
			Success: true,
			Message: "healthy enough!",
		},
	)
}
