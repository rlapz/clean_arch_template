package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rlapz/clean_arch_template/src/model"
	"github.com/rlapz/clean_arch_template/src/usecase"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	log         *logrus.Logger
	usecaseUser *usecase.UserUsecase
}

func NewUserController(log *logrus.Logger, userUsecase *usecase.UserUsecase) *UserController {
	return &UserController{
		log:         log,
		usecaseUser: userUsecase,
	}
}

func (u *UserController) Login(ctx *fiber.Ctx) error {
	req := new(model.UserLoginRequest)
	if err := ctx.BodyParser(req); err != nil {
		u.log.Errorf("Failed to parse request body: %+v", err)
		return fiber.ErrBadRequest
	}

	res, err := u.usecaseUser.Login(ctx.UserContext(), req)
	if err != nil {
		u.log.Errorf("Failed to login: %+v", err)
		return err
	}

	return ctx.JSON(
		model.WebResponse[*model.UserResponse]{
			Success: true,
			Data:    res,
		},
	)
}
