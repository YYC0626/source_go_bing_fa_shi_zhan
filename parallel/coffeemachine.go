package parallel

import "sync"

type CoffeeMachine struct {
	Name       string
	CoffeeName string
	Gopher
	Mlock sync.Mutex
}
