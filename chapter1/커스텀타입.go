package main

import "fmt"

var b int

type hotdog int

var c hotdog

func main() {
	b = 30
	c = 31
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
