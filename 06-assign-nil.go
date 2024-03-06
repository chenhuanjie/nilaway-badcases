package main

import "fmt"

type Container struct {
    children []*Container
}

func main() {
    c := &Container{}
    for i := 0; i < 5; i++ {
        c.children = append(c.children, &Container{})
    }
    c.children[2] = nil
    fmt.Println(c.children[3].children)
}
