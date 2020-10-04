package main

import "fmt"

func main() {

	grade1 := 93
	grade2 := 77
	grade3 := 89
	fmt.Printf("Grades are: %v, %v, %v\n\n", grade1, grade2, grade3)

	/**
	Initialising an array
	*/
	grades := [3]int{93, 77, 89} // grades := [...]int{93, 77, 89}
	fmt.Printf("Grades : %v\n\n", grades)

	/**
	Array Operations
	*/
	var students [3]string
	fmt.Printf("Students : %v\n\n", students)
	students[0] = "Ragul"
	students[1] = "Ravindira"
	students[2] = "John"
	fmt.Printf("Students : %v\n", students)
	fmt.Printf("Students #1: %v\n", students[0])        // access specific element
	fmt.Printf("No of Students: %v\n\n", len(students)) // size of an array

	/**
	Array of Array
	*/
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{1, 0, 1}
	identityMatrix[2] = [3]int{1, 1, 0}
	fmt.Printf("IdentityMatrix: %v\n\n", identityMatrix) // size of an array

	/**
	Copy an Array
	*/
	a := [...]int{1, 2, 3}
	b := a
	b[2] = 6
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("")

	/**
	Slices
	*/
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//s := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // can be initialised as an array also
	s1 := s[:]   // slice of all elements
	s2 := s[3:]  // slice from 4th element to the end
	s3 := s[:6]  // slice of first 6 elements
	s4 := s[3:6] // slice of elements from 4th to 6th element (4th,5th,6th)
	s[5] = 55
	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println("")

	/**
	Append
	*/
	x := []int{1, 2, 3, 4, 5, 6, 7}
	y := append(x[:2], x[3:]...)
	fmt.Println(y)

}
