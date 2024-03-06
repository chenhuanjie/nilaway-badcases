package main

type Client struct {
    m map[string]int
}

func (c *Client) sayhello() {
    c.m["test"] = 15
}

func main() {
    c := &Client{}
    c.sayhello()
}
