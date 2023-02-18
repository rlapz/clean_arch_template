package http

import (
	"fmt"

	delivery "clean_arch_template/delivery"
	service "clean_arch_template/service"
)

type httpEntity1Delivery struct {
	service service.Entity1Service
}

func NewHttpEntity1Delivery(service service.Entity1Service) delivery.Entity1Delivery {
	return &httpEntity1Delivery{
		service: service,
	}
}

/* ----------------------------------------------------------------------- */
func (self *httpEntity1Delivery) Init() {
	fmt.Println("http: Initializing...")
}

func (self *httpEntity1Delivery) Deinit() {
	fmt.Println("http: Cleaning up...")
}

func (self *httpEntity1Delivery) Start() error {
	fmt.Println("http: Starting up...")

	self.somePrivateMethod()
	self.service.Act1("")
	return nil
}

func (self *httpEntity1Delivery) Stop() bool {
	fmt.Println("http: Stopping...")
	return true
}

/* ----------------------------------------------------------------------- */

/* Private funcs/methods */
func (self *httpEntity1Delivery) somePrivateMethod() {
	fmt.Println("http: Hello world! from somePrivateMethod()")
}
