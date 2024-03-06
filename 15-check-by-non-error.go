package main

import "fmt"

func gen(p int) (map[string]int, string) {
    if p > 100 {
        return nil, "p > 100"
    }
    return make(map[string]int), ""
}

func main() {
    m, errMsg := gen(50)
    if errMsg != "" {
        fmt.Println(errMsg)
        return
    }
    m["X"] = 76
}
