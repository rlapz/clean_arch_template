package instance

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/rlapz/clean_arch_template/src/config"
	"github.com/rlapz/clean_arch_template/src/model"
)

func NewFiberApp(config *config.Config) *fiber.App {
	return fiber.New(fiber.Config{
		AppName: config.AppName,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			return ctx.Status(code).JSON(
				model.HttpResponse[any]{
					Success: false,
					Message: err.Error(),
				},
			)
		},
		Prefork: config.Http.IsPrefork,
	})
}
