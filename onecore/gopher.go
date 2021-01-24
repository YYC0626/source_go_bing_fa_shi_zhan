package onecore

import (
	"fmt"
	"time"
)

type Gopher struct {
	Name       string
	Id         int
	CoffeeName string
}

var coffeeMachine = CoffeeMachine{}

func (this *Gopher) MakeCoffee(coffeeName string) {
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
	this.TakeCoffee()
	//step 3
	this.DrinkCoffee()

	//unlock
	coffeeMachine.Mlock.Unlock()
}

func (this *Gopher) TakeCoffee() {
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

func (this *Gopher) DrinkCoffee() {
	if this.CoffeeName != "" {
		fmt.Println("Gopher", this.Id, "Drink Coffee", this.CoffeeName)
		time.Sleep(5 * time.Second)
	} else {
		fmt.Println("Gopher", this.Id, "Has No Coffee To Drink")
	}
}
