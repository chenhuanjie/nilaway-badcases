package main

import "fmt"

func genError(x int) (map[string]int, error) {
    if x > 100 {
        return nil, fmt.Errorf("x > 100")
    }
    return make(map[string]int), nil
}

func ast(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    m, err := genError(42)
    ast(err)
    m["X"] = 1
}
