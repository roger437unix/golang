package main

import "fmt"

type Carro struct {
	Marca  string
	Modelo string
	Ano    int
	Cor    string
}

func main() {
	carro := Carro{
		Marca:  "Toyota",
		Modelo: "Corolla",
		Ano:    2022,
		Cor:    "Prata",
	}

	fmt.Printf("%+v\n", carro)
	carro.andouParaFrente()
	fmt.Println(carro.Marca)
	fmt.Println(carro)
}

// Método para utilizar a struct
func (c *Carro) andouParaFrente() {
	c.Marca = "BMW"
	fmt.Println(c.Marca, "andou...")
}
