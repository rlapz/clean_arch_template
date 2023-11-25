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
		Db:  db,
		Log: log,
	}
}

func (u *UserRepo) FindById(id string) (*entity.User, error) {
	return nil, nil
}

func (u *UserRepo) FindByToken(token string) (*entity.User, error) {
	return nil, nil
}

func (u *UserRepo) Update(user *entity.User) error {
	return nil
}
