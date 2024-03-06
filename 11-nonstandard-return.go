package main

import "fmt"

type Data struct {
    data int
}

func gen(x int) (error, *Data) {
    if x < 0 {
        return fmt.Errorf("x < 0"), nil
    }
    return nil, &Data{data: x}
}

func main() {
    e, d := gen(3)
    if e != nil {
        fmt.Println("got error:", e)
        return
    }
    fmt.Println(d.data)
}
