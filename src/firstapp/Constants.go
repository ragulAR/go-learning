package main

import "fmt"

const b int16 = 27

/**
Enumerated Constant block
*/
const (
	_  = iota // ignore first value
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	EB
	ZB
	YB
)

func main() {

	/**
	Declare Constant
	*/
	const a string = "Ragul"
	const b int16 = 99 // shadowing constant
	const c float32 = 6.99
	const d bool = true

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
	fmt.Println("")

	/**
	Arithmetic Operations
	*/
	const x = 11
	fmt.Printf("%v, %T\n", b+x, b+x)
	fmt.Println("")

	/**
	Enumerated Constant
	*/
	fileSize := 4000000000
	fmt.Printf("%.2fGB", fileSize/GB)
}
