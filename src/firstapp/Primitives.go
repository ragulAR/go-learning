package main

import (
	"fmt"
)

func main() {

	/**
	Declaring boolean value
	*/
	var n bool = true // initially false
	fmt.Printf("%v, %T\n", n, n)

	/**
	Integers - Comparing values
	*/
	i := 1 == 1
	j := 2 == 1
	fmt.Printf("\n%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)

	/**
	Declaring int values
	*/
	k := 33 // signed integer
	fmt.Printf("\n%v, %T\n", k, k)

	var l uint = 22 // unsigned integer
	fmt.Printf("%v, %T\n", l, l)

	/**
	Arithmetic Operations - Integer - I
	*/
	a := 10 // 1010 - converted to binary
	b := 3  // 0011 - converted to binary
	fmt.Println("")
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	fmt.Println("")

	/**
	Arithmetic Operations - Integer - II
	*/
	var c int = 10
	var d int8 = 3
	//fmt.Println(c + d) // error (mismatched types int and int8)
	fmt.Println(c + int(d)) // convert type
	fmt.Println("")

	/**
	Arithmetic Operations - Integer - III (bitwise operators)
	*/
	fmt.Println(a & b)  // AND - 0011 = 2
	fmt.Println(a | b)  // OR - 1011 = 11
	fmt.Println(a ^ b)  // XOR - 1001 = 9
	fmt.Println(a &^ b) // NOT - 0100 = 8
	fmt.Println("")

	/**
	Bit shifting
	*/
	e := 8              // 2^3
	fmt.Println(e << 4) // 2^3 * 2^4 = 2^6
	fmt.Println(e >> 2) // 2^3 / 2^2 = 2^0
	fmt.Println("")

	/**
	Floating point numbers - Declaration
	*/
	var f float64 = 3.14
	f = 13.7e72
	f = 2.1E10
	fmt.Printf("%v, %T\n", f, f)
	fmt.Println("")

	/**
	Arithmetic Operations - Float
	*/
	g := 10.4
	h := 3.6
	fmt.Println(g + h)
	fmt.Println(g - h)
	fmt.Println(g * h)
	fmt.Println(g / h)
	fmt.Printf("%v, %T\n", g+h, g+h) // float64 + float64 = float64
	fmt.Println("")

	/**
	Complex numbers - Declaration
	*/
	var x complex128 = 1 + 2i
	fmt.Printf("%v, %T\n", real(x), real(x))
	fmt.Printf("%v, %T\n", imag(x), imag(x))
	fmt.Println("")

	/**
	Complex function
	*/
	var y complex128 = complex(5, 10)
	fmt.Printf("%v, %T\n", y, y)
	fmt.Println("")

	/**
	String - UTF8 - Declaration
	*/
	var s string = "Hello There"
	fmt.Printf("%v, %T\n", s, s)
	fmt.Println("")

	/**
	String concatenation
	*/
	s2 := "Hello World"
	fmt.Printf("%v, %T\n", s+s2, s+s2)
	fmt.Println("")

	/**
	Rune - Declaration
	*/
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)
}
