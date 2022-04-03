package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(log.Lmsgprefix | log.LstdFlags)

	fmt.Println("Hello, World!")
	// fmt.Println("AQQ!!!")
	// fmt.Println(quote.Go())

	for i := 0; i < 5; i++ {
		callModGreetings("kubus")
	}

	// callModGreetings("")
}

func callModGreetings(name string) {
	var message string
	var err error

	message, err = greetings.Hello(name)

	// If an error was returned, print it to the console and
	// exit the program.
	// If no error was returned, print the returned message
	// to the console.
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)

}
