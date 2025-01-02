/* Primeira função em Go */

package main

import "fmt"

func main() {
	var  num1, num2 float64
	fmt.Print("\nInforme o 1º valor: ")
	fmt.Scanln(&num1)
	fmt.Print("Informe o 2º valor: ")
	fmt.Scanln(&num2)
	fmt.Printf("\n%v + %v = %v\n\n", num1, num2, somar(num1, num2))
}

// Soma dois valores decimais com retorno destes
func somar(n1 float64, n2 float64) float64 {
	soma := n1 + n2
	return soma
}
