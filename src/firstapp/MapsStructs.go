package main

import "fmt"

/**
Struct
*/
type Doctor struct {
	Number     int
	ActorName  string
	Companions []string
}

/**
Inheriting Structs
*/
type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {

	/**
	Initialising a map
	*/
	//statePopulations = make(map[string]int) // using make function
	//statePopulations = map[string]int{
	statePopulations := map[string]int{
		"Colombo": 1234556,
		"Jaffna":  929291,
		"Galle":   193939,
		"Gampaha": 122333,
	}
	fmt.Println(statePopulations, "\n")

	statePopulations["Kalutara"] = 573830 // add new value
	delete(statePopulations, "Galle")     // delete existing value
	fmt.Println(statePopulations, "\n")

	pop, ok := statePopulations["Colombo"] // check presence
	fmt.Println(pop, ok, "\n")

	fmt.Println(len(statePopulations), "\n")

	/**
	Initialising a struct
	*/
	aDoctor := Doctor{
		Number:    3,
		ActorName: "Ragul Ravindira",
		Companions: []string{
			"Messi",
			"Beckham",
		},
	}
	fmt.Println(aDoctor)                     // print all values
	fmt.Println(aDoctor.Companions[1], "\n") // access specific val

	/**
	Initialising an inherited struct
	*/
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)
	fmt.Println(b.Name)
}
