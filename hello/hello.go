package main

import (
	"fmt"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("AQQ!!!")
	fmt.Println(quote.Go())

	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
