package main

import (
	http "clean_arch_template/delivery/http"
	repo "clean_arch_template/repo"
	service "clean_arch_template/service"
)

func main() {
	repo := repo.NewMysqlEntity1Repo()
	service := service.NewEntity1Service(repo)
	delivery := http.NewHttpEntity1Delivery(service)

	delivery.Init()
	defer delivery.Deinit()

	if err := delivery.Start(); err != nil {
		panic(err.Error())
	}
}
