package main

import (
	"fmt"
	"log"

	"github.com/gowithme"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := gowithme.Hellos(names)
	// Request a greeting message.
	message, err := gowithme.Hello("malik")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	// Get a greeting message and print it.
	fmt.Println(message)
	fmt.Println(messages)
}
