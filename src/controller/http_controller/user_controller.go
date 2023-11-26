package http_controller

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
		model.HttpResponse[*model.UserResponse]{
			Success: true,
			Data:    res,
		},
	)
}

// test
func (u *UserController) GetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	ret, err := u.usecaseUser.GetById(ctx.UserContext(), id)
	if err != nil {
		u.log.Errorf("Failed to get user by id: \"%s\":%+v", id, err)
		return err
	}

	if ret == nil {
		return fiber.ErrNotFound
	}

	return ctx.JSON(
		model.HttpResponse[*model.UserResponse]{
			Success: true,
			Data:    ret,
		},
	)
}
