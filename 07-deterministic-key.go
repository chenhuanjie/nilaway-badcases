package main

import "fmt"

type Data struct {
    data int
}

func main() {
    keys := []string{"a", "b", "c"}
    m := make(map[string]*Data)
    for i, k := range keys {
        m[k] = &Data{data: i}
    }
    for _, k := range keys {
        fmt.Println(m[k].data)
    }
}
