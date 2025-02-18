package main

import (
	"fmt"
	connection "go/basics/database"
	"go/basics/modules/greetings"
	"log"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("Abil")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
	names := []string{"Gladys", "Samantha", "Darrin"}
	mess, er := greetings.Hellos(names)
	if er != nil {
		log.Fatal(er)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(mess)
	connection.ConnectDb()
}
