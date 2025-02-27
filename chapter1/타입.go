package main

import "fmt"

// go 는 타입이 전부다. 중심이다.
// "%T\n" 타입을 가져와라 이다 Printf 에서 쓰게되면,
// 정적이기 때문에 미리 정해져있는 타입이 있다.

var y1 = 320
var z1 string = "Shaken, not stirred" // 명시적으로 타입 지정 가능.
// 백쿼터 사용가능
var a string = `James said, "Shaken, not stirred"`

func main() {
	fmt.Println(y1)
	fmt.Printf("%T\n", y1)
	fmt.Println(z1)
	fmt.Printf("%T\n", z1)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
