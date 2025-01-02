package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Informe seu nome: ")
	var nome string
	fmt.Scanln(&nome)
	nome = strings.TrimSpace(nome)
	fmt.Printf("Seja bem-vindo %s!\n", nome)
}
