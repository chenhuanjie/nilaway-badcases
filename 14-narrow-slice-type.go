package main

import "fmt"

type IntArray []int

func (a IntArray) Len() int {
    return len(a) 
}

func main() {
    var x IntArray
    fmt.Println(x.Len())
}
