package main

import "fmt"

func main() {
	x := 54 // 변수를 처음 지정할 때 단축 선언 연산자 사용.  그 다음 부터는 필요없음.
	x = 99

	y := 100 + 224
	fmt.Println(x)
	// go language 에서는 변수를 사용하면 무조건 써야한다. 안그러면 컴파일 오류남.
	fmt.Println(y)

	z := "James bond"
	fmt.Println(z)
}
