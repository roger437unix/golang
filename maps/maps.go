package main

import "fmt"

func main() {
	idade := map[string]int{}
	idade["tux"] = 10
	idade["joca"] = 55
	idade["juca"] = 28
	fmt.Println(idade)
	fmt.Println(idade["tux"])

	fmt.Println()
	anoNasc := map[string]int{
		"maria": 2000,
		"ana":   1977,
	}
	fmt.Println(anoNasc)
	delete(anoNasc, "ana")
	fmt.Println(anoNasc)
}
