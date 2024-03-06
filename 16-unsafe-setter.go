// 这条往往意味着setter方法不安全, 建议修改
package main

import "fmt"

type Data struct {
    data int
}

func insert(m map[string]*Data, key string, value *Data) {
    m[key] = value
}

func main() {
    m := make(map[string]*Data)
    insert(m, "a", &Data{data: 3})
    fmt.Println(m["a"].data)
}
