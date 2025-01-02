package main

import "fmt"

func main() {
	// var nome string = "tux"
	// fmt.Println(nome)
	// nome = "multics"
	// fmt.Println(nome)

	// // Inferência de tipo e atribuição multipla
	// var n1, n2 = 10, 100
	// fmt.Printf("Type: %T -> Value: %v, %v\n", n1, n1, n2)

	// var n3, n4 float64 = 10, 100
	// fmt.Printf("Type: %T -> Value: %v, %v\n", n3, n3, n4)

	// Sem utilizar a palavra reservada "var" com inferência de tipo
	fruta := "apple"
	num := 10.99
	fmt.Printf("Type: %T -> Value: %v\n", fruta, fruta)
	fmt.Printf("Type: %T -> Value: %v\n", num, num)

	// Constantes
	const data_nasc = 1974
	fmt.Printf("Type: %T -> Value: %v\n", data_nasc, data_nasc)
}
