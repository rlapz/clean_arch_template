package service

import (
	"clean_arch_template/entity"
	"clean_arch_template/repo"
)

type entity1Service struct {
	repo repo.Entity1Repo
}

func NewEntity1Service(repo repo.Entity1Repo) Entity1Service {
	return &entity1Service{
		repo: repo,
	}
}

/* ----------------------------------------------------------------------- */
func (self *entity1Service) Act1(Id string) error {
	return nil
}

func (self *entity1Service) Act2(string) (*entity.Entity1, error) {
	return &entity.Entity1{}, nil
}

/* ----------------------------------------------------------------------- */

/* Private funcs/methods */
