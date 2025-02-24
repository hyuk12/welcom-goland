package main

import "fmt"

func main() {
	// 사용하지 않는 변수는 만들지 말아라 _ 를 통해 나타내라.
	n, _ := fmt.Println("Hello, world", 42, true)
	fmt.Println(n)
}
