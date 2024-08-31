package main

import (
	"fmt"
	"log"

	"github.com/peixotoleonardo/golang/tutorials/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("João")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	messages, err := greetings.Hellos([]string{"Gladys", "Samantha", "Darrin"})

	if err != nil {
		log.Fatal(err)
	}

	for _, message = range messages {
		fmt.Println(message)
	}
}
