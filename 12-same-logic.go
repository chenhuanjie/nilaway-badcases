package main

import "fmt"

type Data struct {
    data int
}

func couldBeNil(x int) *Data {
    if x < 0 {
        return nil
    }
    return &Data{data: 2}
}

func main() {
    if tmp := couldBeNil(0); tmp != nil {
        fmt.Println(couldBeNil(0).data)
    }
}
