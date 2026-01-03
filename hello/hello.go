package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Tasuku")
	fmt.Println(message)
}