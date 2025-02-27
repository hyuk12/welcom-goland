package main

import "fmt"

var y3 = 43

func main() {
	fmt.Printf("%#x\n", y3)
	y3 = 911

	fmt.Printf("%#x\t%b\t%x", y3, y3, y3)
	s := fmt.Sprintf("%#x\t%b\t%x", y3, y3, y3)
	fmt.Println(s)
}
