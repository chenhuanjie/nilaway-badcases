package main

func Caller(f func ()) {
    f()
}

func main() {
    var m map[string]int

    Caller(func() {
        m = make(map[string]int)
    })

    m["X"] = 1
}
