package main

import (
	"fmt"
)

func main() {

	myString := "I'm first?"

	defer fmt.Println(myString) // `defer` evaluates here, executes function end

	myString = "Nope, I'm first"

	defer fmt.Println(myString) // `defer` is executed in reverse order, Last First

	// Count down
	for idx := 0; idx < 5; idx++ {
		defer fmt.Println(idx)
	}

	fmt.Println("What is going on?? ")
}
