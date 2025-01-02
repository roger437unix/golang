/* Informe se u m número é par ou impar */

package main

import "fmt"

func main() {
	fmt.Print("\nInforme um número inteiro: ")
	var num int
	fmt.Scanf("%d", &num)
	fmt.Printf("\n%d é %s\n\n", num, parImpar(num))
}


func parImpar(n int) string {
	if n % 2 == 1 {
		return "impar"
	} else {
		return "par"
	}
}
