package usecase

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rlapz/clean_arch_template/src/model"
	"github.com/rlapz/clean_arch_template/src/repo"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	log      *logrus.Logger
	validate *validator.Validate

	userRepo *repo.UserRepo
}

func NewUserUsecase(log *logrus.Logger, validate *validator.Validate, userRepo *repo.UserRepo) *UserUsecase {
	return &UserUsecase{
		log:      log,
		validate: validate,
		userRepo: userRepo,
	}
}

func (u *UserUsecase) Login(ctx context.Context, userReq *model.UserLoginRequest) (*model.UserResponse, error) {
	if err := u.validate.Struct(userReq); err != nil {
		u.log.Errorf("[%s]: Invalid request body: %+v:", userReq.Id, err)
		return nil, fiber.ErrBadRequest
	}

	newUser, err := u.userRepo.FindById(userReq.Id)
	if err != nil {
		u.log.Errorf("[%s]: Failed find user by id: %+v:", userReq.Id, err)
		return nil, fiber.ErrUnauthorized
	}

	if err := bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte(userReq.Password)); err != nil {
		u.log.Errorf("[%s]: Failed compare the password: %+v:", userReq.Id, err)
		return nil, fiber.ErrUnauthorized
	}

	newUser.Token = uuid.New().String()
	if err := u.userRepo.Update(newUser); err != nil {
		u.log.Errorf("[%s]: Failed update user: %+v:", userReq.Id, err)
		return nil, fiber.ErrInternalServerError
	}

	return &model.UserResponse{
		Id:        newUser.Id,
		Name:      newUser.Name,
		Token:     newUser.Token,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}, nil
}
