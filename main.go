package main

import (
	http "clean_arch_template/delivery/http"
	repo "clean_arch_template/repo"
	service "clean_arch_template/service"
)

func main() {
	repo := repo.NewMysqlEntityRepo()
	service := service.NewEntity1Service(repo)
	delivery := http.NewHttpEntity1(service)

	delivery.Init()
	defer delivery.Deinit()

	if err := delivery.Start(); err != nil {
		panic(err.Error())
	}
}
