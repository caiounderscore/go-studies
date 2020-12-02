package main

import (
    "fmt"
)

const (
    _ = iota
    x = 100 //The type is definied when the variable is used
    y = iota * 10
    z = y << iota //byte dislocation
)

func main() {


    s := "oie úoÔÀ" //slice of bytes
    sb := []byte(s)

    fmt.Printf("%v\n%T\n", sb, sb)

    for _, v := range s {
        fmt.Printf("%v - %T - %#U - %#x\n", v, v, v, v)
    }

    fmt.Printf("%v\t%v\t%v", x, y, z)

}
