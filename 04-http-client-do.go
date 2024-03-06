package main

import (
    "bytes"
    "net/http"
    "fmt"
    "io"
)

func main() {
    body := bytes.NewBuffer(nil)
    req, err := http.NewRequest("GET", "https://www.baidu.com", body)
    if err != nil {
        fmt.Println("bad request:", err)
        return
    }
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Println("bad response:", err)
        return
    }
    content, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("read error:", err)
        return
    }
    fmt.Println(string(content))
}
