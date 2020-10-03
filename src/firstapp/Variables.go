package main

import (
	"fmt"
	"strconv"
)

var a float32 = 99

func main() {
	/**
	declaring a variable
	*/
	var i int
	i = 66
	i = 15

	var j int = 66
	i = 25

	k := 96
	k = 9

	fmt.Println("i =", i)
	fmt.Println("j =", j)
	fmt.Println("k =", k)

	fmt.Printf("%v, %T", a, a) // print with types
	fmt.Println("\n")

	/**
	shadowing variable
	*/
	fmt.Printf("%v, %T\n", a, a)
	var a int = 42
	fmt.Printf("%v, %T\n", a, a)

	/**
	declaring acronyms
	*/
	var theURL string = "\nwww.google.com\n "
	fmt.Println(theURL)

	/**
	convert variable type
	*/
	var n int = 11
	fmt.Printf("%v, %T\n", n, n)

	var m float32
	m = float32(n)
	fmt.Printf("%v, %T\n", m, m)

	var o string
	o = strconv.Itoa(n)
	fmt.Printf("%v, %T\n", o, o)
}
