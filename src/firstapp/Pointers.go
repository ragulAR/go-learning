package main

import "fmt"

func main() {

	/**
	Declaring pointers
	*/
	var a int = 44
	var b *int = &a
	fmt.Println(a, b)
	fmt.Println(&a, b) // prints the address location
	fmt.Println(a, *b) // prints the value
	a = 27
	fmt.Println(a, *b)
	*b = 11
	fmt.Println(a, *b)
	fmt.Println("")

	/**
	Pointer Arithmetics
	*/
	i := [3]int{1, 2, 3}
	j := &i[0]
	k := &i[1]
	fmt.Printf("%v %p %p\n\n", i, j, k)
	// use "unsafe" package to perform pointer arithmetics in go

	/**
	Create pointer types
	*/
	// declare a struct
	//var ms myStruct
	//ms = myStruct{foo: 22}
	//fmt.Println(ms)

	// pointer to a struct
	var ms *myStruct
	fmt.Println(ms) // initially value will be initialised to nil val
	//ms = &myStruct{foo: 22}
	ms = new(myStruct) // initialise object with "new" keyword
	fmt.Println(ms)    // ms is holding the address of an object that has a field "foo" of val 42

	ms.foo = 33         // similar to -> (*ms).foo = 33 // dereference a pointer
	fmt.Println(ms.foo) // similar to ->fmt.Println((*ms).foo)
}

type myStruct struct {
	foo int
}
