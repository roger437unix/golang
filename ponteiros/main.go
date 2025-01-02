package main

import "fmt"

func main() {
	var n1, n2 float64 = 8, 5
	var p1 *float64 = &n1
	var p2 *float64 = &n2
	var produto float64 = *p1 * *p2

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(*p1, " x ", *p2)
	fmt.Println(produto)

	*p1, *p2 = 10, 15
	produto = *p1 * *p2
	fmt.Println()
	fmt.Println(n1, " x ", n2)
	fmt.Println(produto)
}
