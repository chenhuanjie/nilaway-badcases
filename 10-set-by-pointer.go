package main

import "fmt"

type Data struct{
    data int
}

func Write(p **Data) {
    *p = &Data{data: 456}
}

func main() {
    var t *Data
    Write(&t)
    fmt.Println(t.data)
}
