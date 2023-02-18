package http

import (
	"fmt"

	delivery "clean_arch_template/delivery"
	service "clean_arch_template/service"
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

	self.somePrivateMethod()
	self.service.Act1("")
	return nil
}

func (self *httpEntity1) Stop() bool {
	fmt.Println("http: Stopping...")
	return true
}

/* ----------------------------------------------------------------------- */

/* Private funcs/methods */
func (self *httpEntity1) somePrivateMethod() {
	fmt.Println("http: Hello world! from somePrivateMethod()")
}
