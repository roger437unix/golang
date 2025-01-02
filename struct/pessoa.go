package main

import "fmt"

type Pessoa struct {
	Nome   string
	Idade  int
	Peso   float64
	Altura float64
}

// Herança
type Jogador struct {
	Pessoa
	Time string
}

func main() {
	p1 := Pessoa{"tux", 20, 50, 1.55}
	fmt.Println(p1)

	p2 := Pessoa{Nome: "multics", Idade: 55}
	fmt.Println(p2)
	fmt.Println(p2.Nome, p2.Idade)

	pessoas := []Pessoa{}
	pessoas = append(pessoas, p1, p2)
	fmt.Println(pessoas)
	fmt.Println(pessoas[0].Nome)

	// map
	fmt.Println()
	alunos := map[string][]Pessoa{}
	alunos["programação"] = pessoas
	fmt.Println(alunos)
	fmt.Println(alunos["programação"])
	fmt.Println(alunos["programação"][0].Nome)

	fmt.Println()
	var professores = map[string][]Pessoa{
		"Geologia": {{Nome: "Tux", Idade: 20, Peso: 50, Altura: .55}, {Nome: "Multics, Idade:55"}},
		"Biologia": {{Nome: "Tux2"}, {Nome: "Multics2"}},
	}
	fmt.Println(professores)
	fmt.Println(professores["Geologia"])
	fmt.Println(professores["Geologia"][0].Nome)

	// Usando struct com herança
	fmt.Println()
	jogador1 := Jogador{p1, "XV de Piracicaba"}
	fmt.Println(jogador1)
	fmt.Println(jogador1.Time)
	fmt.Println(jogador1.Pessoa.Nome)
}
