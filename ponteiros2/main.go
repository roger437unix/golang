package main

import "fmt"

func main() {
	var n1, n2 float64
	fmt.Printf("Digite o 1º número: ")
	fmt.Scan(&n1)
	fmt.Printf("Digite o 2º número: ")
	fmt.Scan(&n2)

	fmt.Println(*somar(&n1, &n2))
}

func somar(pNum1 *float64, pNum2 *float64) *float64 {
	soma := *pNum1 + *pNum2
	return &soma
}

func subtrair(pNum1 *float64, pNum2 *float64) *float64 {
	sub := *pNum1 - *pNum2
	return &sub
}
