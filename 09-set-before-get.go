package main

import "fmt"

type Data struct {
    data int
}

func getMap(key string) *Data {
    m := make(map[string]*Data)

    if m[key] == nil {
        tmp := new(Data)
        tmp.data = 15
        m[key] = tmp
    }

    return m[key]
}

func main() {
    fmt.Println(getMap("nihao").data)
}
