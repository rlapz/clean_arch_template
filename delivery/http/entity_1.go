package http

import (
	delivery "clean_arch_template/delivery"
	"clean_arch_template/service"
	"fmt"
)

type httpEntity1 struct {
	service service.Entity1Service
}

func NewHttpEntity1(service service.Entity1Service) delivery.Entity1Delivery {
	return &httpEntity1{
		service: service,
	}
}

/* ----------------------------------------------------------------------- */
func (self *httpEntity1) Init() {
	fmt.Println("http: Initializing...")
}

func (self *httpEntity1) Deinit() {
	fmt.Println("http: Cleaning up...")
}

func (self *httpEntity1) Start() error {
	fmt.Println("http: Starting up...")
	return nil
}

func (self *httpEntity1) Stop() bool {
	fmt.Println("http: Stopping...")
	return true
}

/* ----------------------------------------------------------------------- */

/* Private funcs/methods */
