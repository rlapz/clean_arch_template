package service

import (
	"fmt"

	entity "clean_arch_template/entity"
	repo "clean_arch_template/repo"
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
	self.somePrivateMethod()
	self.repo.Act1("")
	return nil
}

func (self *entity1Service) Act2(string) (*entity.Entity1, error) {
	return &entity.Entity1{}, nil
}

/* ----------------------------------------------------------------------- */

/* Private funcs/methods */
func (self *entity1Service) somePrivateMethod() {
	fmt.Println("service: Hello world! from somePrivateMethod()")
}
