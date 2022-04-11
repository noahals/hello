package main

import (
	"fmt"
	"log"

	"github.com/noahals/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// get a greetings message and print it
	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
