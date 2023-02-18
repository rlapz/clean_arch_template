package repo

import (
	"fmt"

	entity "clean_arch_template/entity"
)

type mysqlEntity1Repo struct {
}

func NewMysqlEntityRepo() Entity1Repo {
	return &mysqlEntity1Repo{}
}

/* ----------------------------------------------------------------------- */
func (self *mysqlEntity1Repo) Act1(Id string) error {
	self.somePrivateMethod()
	return nil
}

func (self *mysqlEntity1Repo) Act2(Id string) (*entity.Entity1, error) {
	return &entity.Entity1{}, nil
}

/* ----------------------------------------------------------------------- */

/* Private funcs/methods */
func (self *mysqlEntity1Repo) somePrivateMethod() {
	fmt.Println("repo: Hello world! from somePrivateMethod()")
}
