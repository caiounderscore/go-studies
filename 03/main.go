package main

import (
    "fmt"
    "runtime"
)

var x bool
var y rune
var z uint16


func main() {
    z = 65535
    fmt.Println(x, z)
    x = 100 > 10
    y := []rune("AB")
    z++ //cool
    fmt.Println(x,y,z)
    fmt.Println(runtime.GOOS, runtime.GOARCH)
}
