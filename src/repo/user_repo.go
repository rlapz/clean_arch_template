package repo

import (
	"database/sql"

	"github.com/rlapz/clean_arch_template/src/entity"
	"github.com/sirupsen/logrus"
)

type UserRepo struct {
	Db  *sql.DB
	Log *logrus.Logger
}

func NewUserRepo(log *logrus.Logger, db *sql.DB) *UserRepo {
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
