package main

var m map[string]int

// func Init() {
//     m = make(map[string]int)
// }
// 
// func main() {
//     Init()
//     m["X"] = 13
// }

// go提供的init方法也分析不了
func init() {
    m = make(map[string]int)
}

func main() {
    m["X"] = 13
}
