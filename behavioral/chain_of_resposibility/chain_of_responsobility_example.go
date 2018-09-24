package chain_of_responsibility

import "fmt"

func Example() {
	// chain of responsobility example
	handlers := &ConcreteHandlerA{
		Next:&ConcreteHandlerB{
			Next:&ConcreteHandlerC{},
		},
	}

	fmt.Println(handlers.SendRequest(4))
}
