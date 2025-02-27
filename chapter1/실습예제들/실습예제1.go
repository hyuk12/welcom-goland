package main

import "fmt"

// 단축 선언 연산자를 사용해서 이값들을 식별자가 x, y, z 인 변수에 넣는다.
// 단일 프린트 구문, 다중 프린트 구문
/**
1. 42
2. "James Bond"
3. true
*/
var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Printf("%v\n%v\n%v\n", x, y, z)
}
