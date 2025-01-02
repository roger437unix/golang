package main

import "fmt"

func main() {
	var num int = 10
	fmt.Println(num)

	alteraValor(&num)
	fmt.Println(num)
}

func alteraValor(p1 *int) {
	*p1 = 8080
}
