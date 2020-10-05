package main

import "fmt"

func main() {
	fmt.Println("Hello, World!\n")
	printMessage("Hi", "I am Ragul!")
	fmt.Println("")

	firstname := "Ragul"
	lastname := "Ravindira"
	printMessageTwo(&firstname, &lastname) // pointer parameters
	fmt.Println("")

	sum("The sum is: ", 1, 2, 3, 4, 5) // variatic parameters
	fmt.Println("")

	t := total(1, 2, 3, 4, 5)
	fmt.Println("The total is: ", t)
	fmt.Println("")

	ts := totalSum(1, 2, 3, 4, 5)
	fmt.Println("The total sum is: ", *ts)
	fmt.Println("")

	d, err := divide(6.0, 4.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
	fmt.Println("")

	/**
	Anonymous functions
	*/
	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}

	/**
	Anonymous functions as variable
	*/
	f := func() { // var f func() = func() {}
		fmt.Println("\nAnonymous functions as variable")
	}
	f()
	fmt.Println("")

	/**
	Method Invocation
	*/
	g := greeter{
		greeting: "Hello",
		name:     "GO",
	}
	g.greet()
	fmt.Println("The new g.name is: ", g.name)
}

/**
Function with multiple parameter
*/
func printMessage(greeting string, name string) { //func printMessage(greeting, name string) { // if same type
	fmt.Println(greeting, name)
}

/**
Function with pointer parameters
*/
func printMessageTwo(firstname, lastname *string) {
	fmt.Println(*firstname, *lastname)
	fmt.Println(firstname, lastname) // prints the address
	*lastname = "Dravid"
	fmt.Println(*firstname, *lastname)
}

func sum(msg string, values ...int) {

	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(msg, result)
}

/**
Function with return values
*/
func total(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

/**
Function with return values as pointers
*/
func totalSum(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

/**
Function with multiple return values
*/
func divide(a, b float64) (float64, error) {

	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

/**
Method Invocation
*/
type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() { // func (g greeter) greet() { // attributes can be changed when parsed as pointer
	fmt.Println(g.greeting, g.name)
	g.name = ""
}
