/****************************************
Segunda função em Go
- Função que retorna dois valores
- Atribuição multipla em Go
- Ignorando um dos retonos com '_' (underline)
****************************************/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("\nDigite seu primeiro nome: ")
	var frase string
	fmt.Scanf("%s", &frase)

	// Recebendo o retorno com atribuição multipla
	maiusculo, minusculo := mudaTexto(frase)

	fmt.Println("\n" + maiusculo)
	fmt.Println(minusculo + "\n")

	// Ignorando um dos retornos com o underline
	m1, _:= mudaTexto("texto original em minúsculo")

	fmt.Printf("%v\n\n", m1)
}


func mudaTexto(texto string) (string, string) {
	upper := strings.ToUpper(texto)
	lower := strings.ToLower(texto)
	return upper, lower
}
