package delivery

type Delivery interface {
	Init()
	Deinit()
	Start() error
	Stop() bool
}

type Entity1Delivery interface {
	Delivery
}
