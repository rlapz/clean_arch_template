package http_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rlapz/clean_arch_template/src/model"
	"github.com/rlapz/clean_arch_template/src/usecase"
	"github.com/rlapz/clean_arch_template/src/util"
)

type UserController struct {
	log         *util.Logger
	usecaseUser *usecase.UserUsecase
}

func NewUserController(logger *util.Logger, userUsecase *usecase.UserUsecase) *UserController {
	return &UserController{
		log:         logger,
		usecaseUser: userUsecase,
	}
}

func (u *UserController) Login(ctx *fiber.Ctx) error {
	req := new(model.UserLoginRequest)
	if err := ctx.BodyParser(req); err != nil {
		u.log.Errorf("Login: Failed to parse request body: %+v", err)
		return fiber.ErrBadRequest
	}

	res, err := u.usecaseUser.Login(ctx.UserContext(), req)
	if err != nil {
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
