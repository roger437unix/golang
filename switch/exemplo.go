package main

import "fmt"

func main() {
	var num int
	for {
		fmt.Print("Digite um nº inteiro entre 1 e 7: ")
		fmt.Scan((&num))
		if num >= 1 && num <= 7 {
			break
		}
	}
	switch num {
	case 1:
		fmt.Println("Domingo")
	case 2:
		fmt.Println("Segunda")
	case 3:
		fmt.Println("Terça")
	case 4:
		fmt.Println("Quarta")
	case 5:
		fmt.Println("Quinta")
	case 6:
		fmt.Println("Sexta")
	case 7:
		fmt.Println("Sábado")
	default:
		fmt.Println("Número inválido")
	}
}
