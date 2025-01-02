/* Calcula o IMC */

package main

import "fmt"


func main() {
	var peso, altura float64
	fmt.Print("\nInforme o peso: ")
	fmt.Scanf("%f", &peso)
	fmt.Print("Informe a altura: ")
	fmt.Scanln(&altura)

	imc := peso / (altura * altura)

	fmt.Printf("\nPeso: %v\nAltura: %v\nIMC: %.2f\n\n", peso, altura, imc)
}

