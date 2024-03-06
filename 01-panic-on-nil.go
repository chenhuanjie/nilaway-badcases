// 这条实际上是由于nilaway对全局变量分析特殊处理导致的, 如果m由errSource返回则不会报错
package main

import "fmt"

var m map[string]int

func errSource(p int) error {
    if p > 50 {
        return fmt.Errorf("p > 50")
    }
    m = make(map[string]int)
    return nil
}

func main() {
    err := errSource(10)
    if err != nil {
        panic(err)
    }
    m["X"] = 13
}
