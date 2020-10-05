package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	/**
	Defer function
	*/
	fmt.Println("start") // LIFO order
	defer fmt.Println("middle")
	fmt.Println("end")

	a := "hi"
	defer fmt.Println(a)
	a = "hello"

	/**
	Panic function
	*/
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello GO!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Start Panic")
	panicker()
	fmt.Println("End Panic")
}

func panicker() {
	fmt.Println("About to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
			panic(err)
		}
		panic("something bad happened")
		fmt.Println("Panicking done")
	}()
}
