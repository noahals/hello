package main

import (
	"fmt"

	"github.com/noahals/greetings"
)

func main() {
	// get a greetings message and print it
	message := greetings.Hello("Noah")
	fmt.Println(message)
}
