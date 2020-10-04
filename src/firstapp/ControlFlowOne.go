package main

import "fmt"

func main() {

	number := 50
	guess := 30

	/**
	Declaring If Statements
	*/
	if guess < number {
		fmt.Println("Too Low")
	} else if guess > number {
		fmt.Println("Too High")
	} else {
		fmt.Println("Correct Guess!")
	}
	fmt.Println(number <= guess, number >= guess, number != guess) // in-line if conditions

	/**
	Declaring Switch Statements
	*/
	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")
		//fallthrough
	case i <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}

}
