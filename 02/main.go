package main

import (
    "fmt"
)


type sometype int

//package level scope
var b sometype = 10
var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
    fmt.Printf("Values: %v, %v, %v, %v \n", b,x,y,z)
    fmt.Printf("Types: %T, %T, %T, %T \n", b,sometype(x),y,z)
    s := fmt.Sprintf("%v\t%v\t%v\t%v", b,x,y,z)
    fmt.Printf(s)

}
