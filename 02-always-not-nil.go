package main

import "fmt"

func gen() (map[string]int, error) {
    m := make(map[string]int)
    return m, fmt.Errorf("test error")
}

func main() {
    m, _ := gen()
    m["X"] = 13
}
