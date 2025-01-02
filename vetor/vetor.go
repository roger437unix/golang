package main

import "fmt"

func main() {
	var paises [3]string
	paises[0] = "br"
	paises[1] = "it"
	paises[2] = "jp"
	fmt.Println(paises)
	fmt.Println(paises[:])
	fmt.Println(paises[0:2])
	fmt.Println(paises[1:])

	uf := [3]string{"sp", "rj", "mg"}
	fmt.Println(uf)
	fmt.Println()

	// Slices [Tamanho não fixo]
	p2 := []string{}
	fmt.Println(p2)
	p2 = append(p2, "it", "ar", "gb", "ko")
	fmt.Println(p2)
}
