package main

import "fmt"

type Client struct {
	Int
	t int
	String
}

type Int int
type String string

func (c *Int) test1() {
	fmt.Println("test1 was called:", *c)
}

func (s *String) test2() {
	fmt.Println("test2 was called")
}

func main() {
	var client Client
	client.Int = 13
	client.test1()
	client.test2()
}
