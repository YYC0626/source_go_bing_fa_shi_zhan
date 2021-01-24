package parallel

import (
	"fmt"
	"sync"
	"time"
)

type Gopher struct {
	Name       string
	Id         int
	CoffeeName string
}

var coffeeMachineArr = [2]CoffeeMachine{}

func init() {
	var coffeeMachine = CoffeeMachine{Name: "CoffeeMachine1"}
	var coffeeMachine2 = CoffeeMachine{Name: "CoffeeMachine2"}
	coffeeMachineArr[0] = coffeeMachine
	coffeeMachineArr[1] = coffeeMachine2
}

func (this *Gopher) MakeCoffee(coffeeName string, wg *sync.WaitGroup) {

	coffeeMachine := coffeeMachineArr[this.Id%2]

	//lock
	coffeeMachine.Mlock.Lock()

	//step 1
	if coffeeMachine.CoffeeName == "" {
		coffeeMachine.CoffeeName = coffeeName
		coffeeMachine.Gopher = *this
		fmt.Println("Gopher", this.Id, "Make coffee", coffeeMachine.CoffeeName)
		time.Sleep(10 * time.Second)
	}
	//step 2
	this.TakeCoffee(coffeeMachine)
	//step 3
	this.DrinkCoffee(coffeeMachine)

	//unlock
	coffeeMachine.Mlock.Unlock()

	wg.Done()
}

func (this *Gopher) TakeCoffee(coffeeMachine CoffeeMachine) {
	if coffeeMachine.CoffeeName != "" {
		fmt.Println("Gopher", this.Id, "Take Coffee", coffeeMachine.CoffeeName)
		this.CoffeeName = coffeeMachine.CoffeeName
		time.Sleep(5 * time.Second)

		coffeeMachine.CoffeeName = ""
	} else {
		fmt.Println("Gopher", this.Id, "Has No Coffee to Take")
		this.CoffeeName = ""
	}
}

func (this *Gopher) DrinkCoffee(coffeeMachine CoffeeMachine) {
	if this.CoffeeName != "" {
		fmt.Println("Gopher", this.Id, "Drink Coffee", this.CoffeeName)
		time.Sleep(5 * time.Second)
	} else {
		fmt.Println("Gopher", this.Id, "Has No Coffee To Drink")
	}
}
