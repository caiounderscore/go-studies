package main

import (
    "fmt"
)

// Go is a static language

var a bool   //Only Declaration           
var b int
var c string
var d float64
var x = false //Use "var" for package level scope, "var" and ":=" automatically define the type of variable
// a = true //wrong because the variable was declarared before, so i can declarate only on codeblock


func main() {
    a = 0.1 < 20 //Inicialization and Attribution
    fmt.Printf("Value: %v Type: %T \n", a, a)
    a = 20 > 10 // Only Attribution
    someFunction(a,b,c,d)
    someFunction(x,b,c,d)
}

func someFunction(a bool, b int, c string, d float64) {
    fmt.Printf("Original value of a:%v, b:%v, c:%v, d:%f \n", a, b, c, d) //Print Literals, Zero Value in Golang

}
