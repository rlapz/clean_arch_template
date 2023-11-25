package repo

import (
	"database/sql"

	"github.com/rlapz/clean_arch_template/src/entity"
	"go.uber.org/zap"
)

type UserRepo struct {
	Db  *sql.DB
	Log *zap.SugaredLogger
}

func NewUserRepo(log *zap.SugaredLogger, db *sql.DB) *UserRepo {
	return &UserRepo{
		Log: log,
	}
}

func (u *UserRepo) FindById(userRet *entity.User, id string) error {
	_ = u.Db
	return nil
}

func (u *UserRepo) FindByToken(userRet *entity.User, token string) error {
	return nil
}

func (u *UserRepo) Update(userIn *entity.User) error {
	return nil
}
