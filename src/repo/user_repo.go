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
	const query = `
		select  id,
				name,
				password,
				token,
				created_at,
				updated_at
		from user where id = ?
	`

	r, err := u.Db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	ret := new(entity.User)
	if !r.Next() {
		return nil, nil
	}

	if err := r.Scan(&ret.Id, &ret.Name, &ret.Password, &ret.Token, &ret.CreatedAt, &ret.UpdatedAt); err != nil {
		return nil, err
	}

	return ret, nil
}

func (u *UserRepo) FindByToken(token string) (*entity.User, error) {
	return nil, nil
}

func (u *UserRepo) Update(user *entity.User) error {
	return nil
}
