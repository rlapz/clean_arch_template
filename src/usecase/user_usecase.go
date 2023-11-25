package usecase

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rlapz/clean_arch_template/src/entity"
	"github.com/rlapz/clean_arch_template/src/model"
	"github.com/rlapz/clean_arch_template/src/repo"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	log      *zap.SugaredLogger
	validate *validator.Validate

	userRepo *repo.UserRepo
}

func NewUserUsecase(log *zap.SugaredLogger, validate *validator.Validate, userRepo *repo.UserRepo) *UserUsecase {
	return &UserUsecase{
		log:      log,
		validate: validate,
		userRepo: userRepo,
	}
}

func (u *UserUsecase) Login(ctx context.Context, userReq *model.UserLoginRequest) (*model.UserResponse, error) {
	tx, err := u.userRepo.Db.BeginTx(ctx, nil)
	if err != nil {
		u.log.Errorf("[%s]: Failed begin transaction: %+v:", userReq.Id, err)
		return nil, fiber.ErrInternalServerError
	}
	defer tx.Rollback()

	if err := u.validate.Struct(userReq); err != nil {
		u.log.Errorf("[%s]: Invalid request body: %+v:", userReq.Id, err)
		return nil, fiber.ErrBadRequest
	}

	newUser := new(entity.User)
	if err := u.userRepo.FindById(newUser, userReq.Id); err != nil {
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

	if err := tx.Commit(); err != nil {
		u.log.Errorf("[%s]: Failed commit transaction: %+v:", userReq.Id, err)
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
