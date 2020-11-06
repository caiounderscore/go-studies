package main

import (
	"fmt"
	"reflect"
)

func main() {
	numBytes, err := fmt.Println("Hello World")
	fmt.Println(numBytes, err)

	//gopher (:=) works only inside code blocks e and define automatically variable types
	a, b, c, d := 0, -1.0, "string", true
	fmt.Println("Golang primitive types:", reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c), reflect.TypeOf(c))

	// = use to insert new values on variables already existences
	a, b, c, d = -1, 50, "opa", false
	fmt.Printf("Golang primitive types: %T, %T, %T, %T", a, b, c, d)
}
