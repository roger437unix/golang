package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("\nDigite a conta [ex: 8*5]: ")
	var strConta string
	fmt.Scan(&strConta)
	conta := strings.Split(strConta, "")

	ponteiro := calcular(conta)
	fmt.Printf("\n%s %s %s = %d\n\n", conta[0], conta[1], conta[2], *ponteiro)
}

func calcular(conta []string) *int {
	n1, _ := strconv.Atoi(conta[0])
	n2, _ := strconv.Atoi(conta[2])
	op := conta[1]
	var result int

	switch op {
	case "+":
		result = n1 + n2
	case "-":
		result = n1 - n2
	case "*":
		result = n1 * n2
	case "/":
		result = n1 / n2
	default:
		// panic(("\n*** Operação inválida ***\n"))
		fmt.Println("*** Operação inválida ***")
	}
	return &result
}
