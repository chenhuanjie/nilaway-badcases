package main

import "fmt"

func genMap(p int) (map[string]int, error) {
    if p < 100 {
        return nil, fmt.Errorf("p < 100")
    }
    return make(map[string]int), nil
}

func replaceError(e error) (error, bool) {
    switch {
    case e != nil:
        return fmt.Errorf("another error: %s", e.Error()), true
    default:
        return nil, false
    }
}

func main() {
    m, e := genMap(155)
    e, valid := replaceError(e)
    if valid {
        fmt.Println("we have a trouble now")
    } else {
        m["X"] = 13
    }
}
