package onecore

import "sync"

//共享资源
type CoffeeMachine struct {
	CoffeeName string
	Gopher
	Mlock sync.Mutex
}
