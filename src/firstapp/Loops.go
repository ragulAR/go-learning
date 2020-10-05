package main

import "fmt"

func main() {

	/**
	Initialising a for loop
	*/
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("")

	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}
	fmt.Println("")

	m := 0
	for ; m < 5; m++ {
		fmt.Println(m)
	}
	fmt.Println("")

	n := 0
	for n < 5 {
		fmt.Println(n)
		n++
	}
	fmt.Println("")

	/**
	Breaking a loop
	*/
	x := 0
	for {
		fmt.Println(x)
		x++
		if x == 6 {
			break // break the loop based on condition
		}
	}
	fmt.Println("")

	/**
	  Nested Loops
	*/
	for i := 1; i < 4; i++ {
		for j := 1; j < 4; j++ {
			fmt.Println(i * j)
		}
	}
	fmt.Println("")

	/**
	  Break Outer Loops
	*/
	Loop:
	for i := 1; i < 4; i++ {
		for j := 1; j < 4; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}
	fmt.Println("")

	/**
	Collections with loops
	*/
	//s := [3]int{1, 2, 3} // int
	s := map[string]int{ // map
		"Colombo": 102093910,
		"Jaffna":  202020,
		"Galle":   20202,
	}
	//for _, v := range s { // get only values
	for k, v := range s {
		//fmt.Println(v)
		fmt.Println(k, v)
	}
	fmt.Println("")

}
