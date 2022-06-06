package main

import (
	"fmt"
	"greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("Gladys")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	names := []string{"Bob", "James", "Michael", "Matthew"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
