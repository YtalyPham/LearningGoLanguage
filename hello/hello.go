package main

import (
	"fmt"
	"log"

	//"rsc.io/quote"

	"example.com/greetings"
)

func main() {
	//fmt.Println(quote.Go())

	//get a greeting message and print it

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Ytaly", "Valentine", "Luna"}

	// Request a greeting message.
	message, err := greetings.Hellos(names)
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)

}
