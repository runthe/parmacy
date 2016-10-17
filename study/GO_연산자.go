package main

import "fmt"

func main() {
	var k int = 10
	var p = &k //k의 주소 할당

	fmt.Println(*p) //p가 가리키는 주소의 값을 출력
}