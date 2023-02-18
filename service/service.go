package service

import (
	entity "clean_arch_template/entity"
)

type Entity1Service interface {
	Act1(string) error
	Act2(string) (*entity.Entity1, error)
}
