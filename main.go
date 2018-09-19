package main

import (
	"fmt"
	"github.com/NickTaporuk/golang_patterns/behavioral/chain_of_resposibility"
)

func main() {

	handlers := &chain_of_responsibility.ConcreteHandlerA{
		Next:&chain_of_responsibility.ConcreteHandlerB{
			Next:&chain_of_responsibility.ConcreteHandlerC{},
		},
	}

	fmt.Println(handlers.SendRequest(4))
}
