package main

import "fmt"

func main() {
    var a, b []int

    for i := 1; i < 10; i++ {
        a = append(a, i)
        b = append(b, i)
    }

    if len(a) > 5 {
        fmt.Println(b[4])
    }
}
