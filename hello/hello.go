package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("PIYO")
	fmt.Printf(message)
}
