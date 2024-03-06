// 这个不算很好, 参数可能会变, nil可能会发生
package main

import "fmt"

type Dummy struct {
    data int
}

func genNil(x int) *Dummy {
    if x == 0 {
        return nil
    }
    return &Dummy{data: x}
}

func targetFunc() *Dummy {
    return genNil(1)
}

func main() {
    d := targetFunc()
    fmt.Println(d.data)
}
